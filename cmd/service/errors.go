package main

import (
	"errors"
	goruntime "runtime"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stevepartridge/versions"
)

var (
	ErrMissingVersionId        = errors.New("version.missing.id")
	ErrVersionInvalidVersionId = errors.New("version.invalid_id")

	ErrVersionCreateMissingVersion = errors.New("version.create.missing.version")

	ErrGetLatestVersionMissingApplicationId = errors.New("version.create.missing.application_id")

	// ErrVersionDeleteMissingId = errors.New("version.delete.invalid.version_id")

	ErrMissingApplicationId            = errors.New("application.missing.id")
	ErrApplicationInvalidApplicationId = errors.New("application.invalid_id")
	ErrRequestMissingApplication       = errors.New("application.request.missing.application")
	ErrRequestMissingApplicationId     = errors.New("application.request.missing.application_id")

	ErrRequestMissingDownload    = errors.New("download.request.missing.download")
	ErrRequestMissingDownloadId  = errors.New("download.missing.id")
	ErrDownloadInvalidDownloadId = errors.New("download.invalid_id")
)

func handleError(err error) error {

	code := statusFromError(err)

	return status.Error(code, err.Error())
}

func statusFromError(err error) codes.Code {

	code := codes.Unknown

	switch err {
	case ErrVersionCreateMissingVersion,
		ErrGetLatestVersionMissingApplicationId,
		ErrMissingApplicationId,
		ErrRequestMissingApplication,
		ErrRequestMissingApplicationId,

		versions.ErrVersionExists,
		versions.ErrVersionIsNil,
		versions.ErrVersionMissingOS,
		versions.ErrVersionMissingArch,

		versions.ErrVersionInvalidId,
		versions.ErrVersionMissingId,
		versions.ErrVersionIdNotEmpty,
		versions.ErrVersionMissingApplication,
		versions.ErrVersionMissingApplicationId,
		versions.ErrVersionMissingComponents,

		versions.ErrApplicationExists,
		versions.ErrApplicationMissingId,

		versions.ErrStorageInvalidType,
		versions.ErrDownloadExists,
		versions.ErrDownloadStorageTypeMissing,
		versions.ErrDownloadDataIsNil,
		versions.ErrDownloadFilenameMissing,
		versions.ErrDownloadIdMissing,
		versions.ErrDownloadVersionIdMissing,

		ErrVersionInvalidVersionId:

		code = codes.InvalidArgument

	case versions.ErrVersionNotFound,
		versions.ErrApplicationNotFound,
		versions.ErrDownloadNotFound,
		versions.ErrDownloadFileNotFound:

		code = codes.NotFound
	}

	return code

}

func httpStatusFromError(err error) int {
	return runtime.HTTPStatusFromCode(statusFromError(err))
}

func ifError(err error) {
	_, file, line, _ := goruntime.Caller(5)
	if err != nil {
		log.Error().
			Str("file", file).
			Int("line", line).
			Err(err).
			Msg("Error")
	}
}
