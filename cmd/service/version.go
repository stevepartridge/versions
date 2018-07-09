package main

import (
	// "fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"

	"github.com/stevepartridge/versions"
	pb "github.com/stevepartridge/versions/protos"
)

func (v *versionService) GetVersion(c context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {

	if req == nil {
		return nil, handleError(ErrVersionCreateMissingVersion)
	}

	if req.VersionId == 0 {
		return nil, handleError(ErrVersionCreateMissingVersion)
	}

	log.Debug().Msgf("payload %+v", req)

	version, err := versionStore.GetById(int(req.VersionId))
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	return &pb.VersionResponse{
		Version: version.ToProto(),
	}, nil

}

func (v *versionService) GetLatestVersion(c context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {

	if req == nil {
		return nil, handleError(ErrVersionCreateMissingVersion)
	}

	if req.ApplicationId == "" {
		return nil, handleError(ErrGetLatestVersionMissingApplicationId)
	}

	log.Debug().Msgf("payload %+v", req)

	version, err := versionStore.GetLatest(req.ApplicationId)
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	return &pb.VersionResponse{
		Version: version.ToProto(),
	}, nil

}

func (v *versionService) GetVersions(c context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {

	if req == nil {
		return nil, handleError(ErrVersionCreateMissingVersion)
	}

	if req.ApplicationId == "" {
		return nil, handleError(ErrGetLatestVersionMissingApplicationId)
	}

	list, err := versionStore.GetByApplication(
		req.ApplicationId,
		int(req.Limit),
		int(req.Offset),
	)

	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	resp := &pb.VersionResponse{
		Versions: versions.VersionsToProtos(list),
		Limit:    req.Limit,
		Offset:   req.Offset,
	}

	return resp, nil

}

func (v *versionService) CreateVersion(c context.Context, req *pb.Version) (*pb.VersionResponse, error) {

	if req == nil {
		return nil, handleError(ErrVersionCreateMissingVersion)
	}

	ver := versions.FromProto(req)

	if req.Application != nil {

		app := versions.ApplicationFromProto(req.Application)

		err := versionStore.CreateApplication(&app)
		if err != nil {
			return nil, handleError(err)
		}

		if app.Id != "" {
			ver.ApplicationId = app.Id
		}

	}

	err := versionStore.Create(&ver)
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	resp := &pb.VersionResponse{
		Version: ver.ToProto(),
	}

	return resp, nil

}

func (v *versionService) UpdateVersion(c context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {

	if req.Version == nil {
		return nil, handleError(ErrVersionCreateMissingVersion)
	}

	if req.VersionId > 0 && req.Version.Id == 0 {
		req.Version.Id = req.VersionId
	}

	ver := versions.FromProto(req.Version)

	if req.Version.Application != nil {

		app := versions.ApplicationFromProto(req.Version.Application)

		err := versionStore.CreateApplication(&app)
		if err != nil {
			return nil, handleError(err)
		}

		if app.Id != "" {
			ver.ApplicationId = app.Id
		}

	}

	err := versionStore.Update(&ver)
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.VersionResponse{
		Version: ver.ToProto(),
	}, nil

}

func (v *versionService) DeleteVersion(c context.Context, req *pb.VersionRequest) (*empty.Empty, error) {

	log.Debug().Msgf("payload %+v", req)

	if req.VersionId == 0 {
		return nil, handleError(ErrMissingVersionId)
	}

	ver, err := versionStore.GetById(int(req.VersionId))
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	if ver.Id == 0 {
		return nil, handleError(ErrVersionInvalidVersionId)
	}

	err = versionStore.Delete(ver.Id)
	ifError(err)

	return &empty.Empty{}, err

}
