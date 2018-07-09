package client

import (
	"crypto/x509"
	// "errors"
	"fmt"
	// "io/ioutil"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	pb "github.com/stevepartridge/versions/protos"
)

type Client struct {
	serviceURI string

	Conn     *grpc.ClientConn
	CertPool *x509.CertPool

	service pb.VersionsClient
}

func NewClient(serviceURI string) (*Client, error) {
	c := &Client{
		serviceURI: serviceURI,
	}

	return c, nil
}

func (c *Client) Connect() error {

	var err error

	if c.CertPool != nil {

		creds := credentials.NewClientTLSFromCert(c.CertPool, "")

		c.Conn, err = grpc.Dial(c.serviceURI, grpc.WithTransportCredentials(creds))
		if err != nil {
			fmt.Println("Error dialing service", err.Error())
			return err
		}

	}

	if c.Conn == nil {
		c.Conn, err = grpc.Dial(c.serviceURI)
		if err != nil {
			fmt.Println("Error", err.Error())
			return err
		}
	}

	if c.service == nil {
		c.service = pb.NewVersionsClient(c.Conn)
	}

	return nil
}

func (c *Client) Check() {

	if c.Conn == nil {
		c.Connect()
	}

	switch c.Conn.GetState() {

	case connectivity.TransientFailure, connectivity.Shutdown:
		err := c.Connect()
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
	default:
		break
	}

	return
}

func (c *Client) Info() (*pb.ServiceInfoResponse, error) {
	return c.service.Info(context.Background(), &pb.ServiceInfoRequest{})
}
