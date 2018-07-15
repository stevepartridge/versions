package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/stevepartridge/env"
	_service "github.com/stevepartridge/service"
	"github.com/stevepartridge/versions"
	pb "github.com/stevepartridge/versions/protos"
)

var (
	serviceName = "versions"
	version     = "0.0.0"
	builtAt     = ""
	build       = "0"

	defaultHost = "versions.local"
	defaultPort = 8000
	host        string
	port        int

	enableInsecure = false

	service     *_service.Service
	grpcService *versionService

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

	localStorage, err := versions.NewLocalStorage()
	ifError(err)

	err = versions.AddStorage(versions.StorageTypeLocal, localStorage)
	ifError(err)

	host = env.Get("HOST")
	if host == "" {
		host = defaultHost
	}

	port = env.GetAsInt("PORT")
	if port == 0 {
		port = defaultPort
	}

	service, err = _service.New(host, port)
	ifError(err)

	serve()

}

func serve() {

	var err error

	enableInsecure = env.GetAsBool("ENABLE_INSECURE")

	if enableInsecure {
		log.Warn().Msg("Insecure is enabled.  Not recommended when in production.")
		service.EnableInsecure()
	}

	rootCA := getRootCA()
	if rootCA != nil {
		err = service.AppendCertsFromPEM(rootCA)
		ifError(err)
	}

	cert, err := getCert()
	ifError(err)
	key, err := getKey()
	ifError(err)

	err = service.AddKeyPair(cert, key)
	ifError(err)

	service.Grpc.AddInterceptors(
		RequestInterceptor(),
		TelemetryInterceptor(),
	)

	grpcService = &versionService{}

	pb.RegisterVersionsServer(service.GrpcServer(), grpcService)

	err = service.EnableGatewayHandler(pb.RegisterVersionsHandlerFromEndpoint)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	httpMiddleware()

	serveSwagger()

	handleFallbackDownloads()

	log.Info().
		Int("port", service.Port).
		Str("host", service.Host).
		Msg("Listening")

	err = service.Serve()
	if err != nil {
		log.Fatal().Err(err).Msg("Listen and Serve")
	}

}
