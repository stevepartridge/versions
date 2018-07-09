package main

import (
	// "fmt"

	"github.com/golang/protobuf/ptypes/empty"
	// "github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/stevepartridge/versions"
	pb "github.com/stevepartridge/versions/protos"
)

func (v *versionService) GetApplication(c context.Context, req *pb.ApplicationRequest) (*pb.ApplicationResponse, error) {

	if req == nil {
		return nil, handleError(ErrRequestMissingApplication)
	}

	if req.ApplicationId == "" {
		return nil, handleError(ErrRequestMissingApplication)
	}

	application, err := versionStore.GetApplicationById(req.ApplicationId)
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	return &pb.ApplicationResponse{
		Application: application.ToProto(),
	}, nil

}

func (v *versionService) GetApplications(c context.Context, req *pb.ApplicationRequest) (*pb.ApplicationResponse, error) {

	if req == nil {
		return nil, handleError(ErrRequestMissingApplication)
	}

	list, err := versionStore.GetApplications(
		int(req.Limit),
		int(req.Offset),
	)

	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	resp := &pb.ApplicationResponse{
		Applications: versions.ApplicationsToProtos(list),
		Limit:        req.Limit,
		Offset:       req.Offset,
	}

	return resp, nil

}

func (v *versionService) CreateApplication(c context.Context, req *pb.Application) (*pb.ApplicationResponse, error) {

	if req == nil {
		return nil, handleError(ErrRequestMissingApplication)
	}

	app := versions.ApplicationFromProto(req)

	err := versionStore.CreateApplication(&app)

	resp := &pb.ApplicationResponse{}

	if err != nil {
		return nil, handleError(err)
	}

	resp.Application = app.ToProto()

	return resp, nil

}

func (v *versionService) UpdateApplication(c context.Context, req *pb.ApplicationRequest) (*pb.ApplicationResponse, error) {

	if req.Application == nil {
		return nil, handleError(ErrRequestMissingApplication)
	}

	app := versions.ApplicationFromProto(req.Application)

	err := versionStore.UpdateApplication(&app)
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.ApplicationResponse{
		Application: app.ToProto(),
	}, nil

}

func (v *versionService) DeleteApplication(c context.Context, req *pb.ApplicationRequest) (*empty.Empty, error) {

	if req.ApplicationId == "" {
		return nil, handleError(ErrRequestMissingApplicationId)
	}

	app, err := versionStore.GetApplicationById(req.ApplicationId)
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	if app.Id == "" {
		return nil, handleError(ErrApplicationInvalidApplicationId)
	}

	err = versionStore.DeleteApplication(app.Id)
	ifError(err)

	return &empty.Empty{}, err

}
