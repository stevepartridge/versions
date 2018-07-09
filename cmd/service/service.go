package main

import (
	// "fmt"

	"golang.org/x/net/context"

	pb "github.com/stevepartridge/versions/protos"
)

type versionService struct{}

func (v *versionService) Info(c context.Context, req *pb.ServiceInfoRequest) (*pb.ServiceInfoResponse, error) {

	return &pb.ServiceInfoResponse{
		BuiltAt: builtAt,
		Version: version,
		Build:   build,
	}, nil

}
