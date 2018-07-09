package client

import (
	"golang.org/x/net/context"

	"github.com/stevepartridge/versions"
	pb "github.com/stevepartridge/versions/protos"
)

// Get Version
// Create Version
// Update Version
// Delete Version
// Get By App?

func (c *Client) CreateVersion(version, name, appId string) (*pb.VersionResponse, error) {

	major, minor, revision := versions.Parse(version)

	return c.service.CreateVersion(context.Background(), &pb.Version{
		Major:         major,
		Minor:         minor,
		Revision:      revision,
		Name:          name,
		ApplicationId: appId,
	})
}

func (c *Client) GetVersion(versionId int) (*pb.VersionResponse, error) {
	return c.service.GetVersion(context.Background(), &pb.VersionRequest{
		VersionId: int32(versionId),
	})
}

func (c *Client) UpdateVersion(versionId int, version, name, appId string, stable bool) (*pb.VersionResponse, error) {
	return c.service.GetVersion(context.Background(), &pb.VersionRequest{
		VersionId: int32(versionId),
	})
}

func (c *Client) DeleteVersion(versionId int) (*pb.VersionResponse, error) {
	return c.service.GetVersion(context.Background(), &pb.VersionRequest{
		VersionId: int32(versionId),
	})
}
