package main

import (
	"math"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/grpc-ecosystem/go-grpc-middleware"
)

var (
	maxReceiveMessageSize = math.MaxInt32
	maxSendMessageSize    = math.MaxInt32

	server *grpc.Server
)

func newGrpcServer() *grpc.Server {

	interceptors := []grpc.UnaryServerInterceptor{
		RequestInterceptor(),
		TelemetryInterceptor(),
	}

	options := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxReceiveMessageSize),
		grpc.MaxSendMsgSize(maxSendMessageSize),
		grpc.Creds(credentials.NewClientTLSFromCert(certPool, host)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptors...,
		)),
	}

	return grpc.NewServer(options...)
}
