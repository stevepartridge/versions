package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	// "fmt"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	"github.com/stevepartridge/versions"
	pb "github.com/stevepartridge/versions/protos"
)

// 1MB
const MAX_MEMORY = 1 * 1024 * 1024

func handleFallbackDownloads() {

	service.Router.Post("/v1/versions/{version_id}/downloads", handleCreateDownload)
	service.Router.Get("/v1/downloads/{download_id}/file", handleGetDownloadFile)

}

func (v *versionService) GetDownload(c context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {

	if req == nil {
		return nil, handleError(ErrRequestMissingDownload)
	}

	if req.DownloadId == 0 {
		return nil, handleError(ErrRequestMissingDownload)
	}

	download, err := versionStore.GetDownloadById(int(req.DownloadId))
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	return &pb.DownloadResponse{
		Download: download.ToProto(),
	}, nil

}

func (v *versionService) GetDownloads(c context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {

	if req == nil {
		return nil, handleError(ErrRequestMissingDownload)
	}

	list, err := versionStore.GetDownloads(
		int(req.Limit),
		int(req.Offset),
	)

	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	resp := &pb.DownloadResponse{
		Downloads: versions.DownloadsToProtos(list),
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	return resp, nil

}

func (v *versionService) CreateDownload(c context.Context, req *pb.Download) (*pb.DownloadResponse, error) {

	if req == nil {
		return nil, handleError(ErrRequestMissingDownload)
	}

	dl := versions.DownloadFromProto(req)

	err := versionStore.CreateDownload(&dl)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.DownloadResponse{}

	if err != nil {
		return nil, handleError(err)
	}

	resp.Download = dl.ToProto()

	return resp, nil

}

// handleCreateDownload Simplifies the usage when making an HTTP 1 request
// TODO: clean up usage of error handling
func handleCreateDownload(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	versionId, _ := strconv.Atoi(chi.URLParam(r, "version_id"))

	dl := &pb.Download{
		VersionId:   int32(versionId),
		StorageType: r.FormValue("storage_type"),
		Os:          r.FormValue("os"),
		Arch:        r.FormValue("arch"),
		Filename:    r.FormValue("filename"),
		Extension:   r.FormValue("extension"),
	}

	// fmt.Println("storage type", dl.VersionId, dl.StorageType)

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			buf, _ := ioutil.ReadAll(file)
			dl.Data = buf
		}
	}

	resp, err := grpcService.CreateDownload(context.Background(), dl)
	if err != nil {
		http.Error(w, err.Error(), httpStatusFromError(err))
		return
	}

	if resp == nil {
		http.Error(w, "Response is nil", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), httpStatusFromError(err))
		return
	}

	w.Write(data)
}

func handleGetDownloadFile(w http.ResponseWriter, r *http.Request) {

	downloadId, _ := strconv.Atoi(chi.URLParam(r, "download_id"))

	download, err := versionStore.GetDownloadById(downloadId)
	if err != nil {
		http.Error(w, err.Error(), httpStatusFromError(err))
		return
	}

	file, err := download.GetFile()
	if err != nil {
		http.Error(w, err.Error(), httpStatusFromError(err))
		return
	}

	http.ServeContent(w, r, download.Filename, download.UpdatedAt, bytes.NewReader(file))

}

func (v *versionService) UpdateDownload(c context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {

	if req.Download == nil {
		return nil, handleError(ErrRequestMissingDownload)
	}

	app := versions.DownloadFromProto(req.Download)

	err := versionStore.UpdateDownload(&app)
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.DownloadResponse{
		Download: app.ToProto(),
	}, nil

}

func (v *versionService) DeleteDownload(c context.Context, req *pb.DownloadRequest) (*empty.Empty, error) {

	if req.DownloadId == 0 {
		return nil, handleError(ErrRequestMissingDownloadId)
	}

	app, err := versionStore.GetDownloadById(int(req.DownloadId))
	if err != nil {
		ifError(err)
		return nil, handleError(err)
	}

	if app.Id == 0 {
		return nil, handleError(ErrDownloadInvalidDownloadId)
	}

	err = versionStore.DeleteDownload(app.Id)
	ifError(err)

	return &empty.Empty{}, err

}
