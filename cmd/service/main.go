package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/stevepartridge/env"
	"github.com/stevepartridge/versions"
	"github.com/stevepartridge/versions/insecure"
	pb "github.com/stevepartridge/versions/protos"
)

var (
	service = "versions"
	version = "0.0.0"
	builtAt = ""
	build   = "0"

	defaultHost = "versions.local"
	defaultPort = 8000
	host        string
	port        int

	enableInsecure = false

	certPool *x509.CertPool

	versionStore versions.Store
)

func main() {

	if builtAt == "" {
		builtAt = time.Now().Format(time.RFC3339Nano)
	}

	logLevel := zerolog.InfoLevel
	switch strings.ToLower(env.Get("LOG_LEVEL")) {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "fatal", "panic":
		logLevel = zerolog.PanicLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	if strings.ToLower(env.Get("LOG_FORMAT")) != "json" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano

	var err error

	versionStore, err = versions.NewStore("")
	ifError(err)

	host = env.Get("HOST")
	if host == "" {
		host = defaultHost
	}

	port = env.GetAsInt("PORT")
	if port == 0 {
		port = defaultPort
	}

	serve()

}

func serve() {

	enableInsecure = env.GetAsBool("ENABLE_INSECURE")

	if enableInsecure {
		log.Warn().Msg("Insecure is enabled.  Not recommended when in production.")
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM([]byte(insecure.RootCA))
	if !ok {
		panic("Invalid or bad certs")
	}

	// Set up the gRPC service
	grpcServer := newGrpcServer()
	pb.RegisterVersionsServer(grpcServer, &versionService{})

	creds := credentials.NewTLS(&tls.Config{
		ServerName:         host,
		RootCAs:            certPool,
		InsecureSkipVerify: enableInsecure,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	// Address for reverse proxy connection (http 1.x requests)
	addr := fmt.Sprintf("%s:%d", host, port)

	// Setup gateway http/s fallback support for REST clients
	gwMux := runtime.NewServeMux()
	err := pb.RegisterVersionsHandlerFromEndpoint(context.Background(), gwMux, addr, opts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	server := newServer(gwMux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	cert, err := getCertificate()
	ifError(err)
	tlsConfig := tls.Config{
		Certificates:       []tls.Certificate{cert},
		NextProtos:         []string{"h2"},
		InsecureSkipVerify: enableInsecure,
	}

	tlsConfig.BuildNameToCertificate()

	srv := &http.Server{
		Addr:      strconv.Itoa(port),
		Handler:   handlerFunc(grpcServer, server),
		TLSConfig: &tlsConfig,
	}

	log.Info().
		Int("port", port).
		Str("host", host).
		Msg("Listening")

	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))

	if err != nil {
		log.Fatal().Err(err).Msg("Listen and Serve")
	}

}

func handlerFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}
