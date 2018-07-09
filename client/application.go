package client

import (
	"golang.org/x/net/context"

	pb "github.com/stevepartridge/versions/protos"
)

// Get Application
// Create Application
// Update Application
// Delete Application

func (c *Client) GetApplication(id string) (*pb.ApplicationResponse, error) {

	return c.service.GetApplication(context.Background(), &pb.ApplicationRequest{
		ApplicationId: id,
	})
}

func (c *Client) CreateApplication(id, name string) (*pb.ApplicationResponse, error) {

	return c.service.CreateApplication(context.Background(), &pb.Application{
		Id:   id,
		Name: name,
	})
}
