package client

import (
	"golang.org/x/net/context"

	pb "github.com/stevepartridge/versions/protos"
)

// Get Download
// Create Download
// Update Download
// Delete Download

func (c *Client) GetDownload(id int) (*pb.DownloadResponse, error) {

	return c.service.GetDownload(context.Background(), &pb.DownloadRequest{
		DownloadId: int32(id),
	})
}

func (c *Client) GetDownloadFile(id int) (*pb.Download, error) {

	return c.service.GetDownloadFile(context.Background(), &pb.Download{
		Id: int32(id),
	})
}

func (c *Client) CreateDownload(versionId int, req *pb.Download) (*pb.DownloadResponse, error) {

	return c.service.CreateDownload(context.Background(), req)
}
