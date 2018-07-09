package main

import (
	// "fmt"
	// "net/http"
	"time"

	"golang.org/x/net/context" // have to use this context because of grpc lib
	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	// "google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"
)

func RequestInterceptor() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		log.Info().
			Str("request_method", info.FullMethod).
			Str("user_agent", getUserAgentFromContext(ctx)).
			Str("ip", getForwardedForFromContext(ctx)).
			Str("referer", getForwardedHostFromContext(ctx)).
			Msg("")

		resp, err := handler(ctx, req)
		ifError(err)

		return resp, err
	}
}

func TelemetryInterceptor() func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		started := time.Now().UTC()

		resp, err := handler(ctx, req)
		ifError(err)

		finished := time.Now().UTC()

		log.Info().
			Str("request_method", info.FullMethod).
			Str("start", started.Format(time.RFC3339Nano)).
			Str("finish", finished.Format(time.RFC3339Nano)).
			Str("duration", finished.Sub(started).String()).
			Str("duration", finished.Sub(started).String()).
			Msg("")

		return resp, err
	}
}

func getUserAgentFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["grpcgateway-user-agent"]; ok && len(key) > 0 {
			return key[0]
		}

		if key, ok := md["user-agent"]; ok && len(key) > 0 {
			return key[0]
		}
	}

	return ""
}

func getForwardedForFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["x-forwarded-for"]; ok && len(key) > 0 {
			return key[0]
		}
	}

	return ""
}

func getForwardedHostFromContext(ctx context.Context) string {

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if key, ok := md["x-forwarded-host"]; ok && len(key) > 0 {
			return key[0]
		}
	}
	return ""
}
