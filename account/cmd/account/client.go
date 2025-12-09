package account

import (
	"context"
	"errors"

	"github.com/olujimiAdebakin/ProtoGraph/account/pb"
	"google.golang.org/grpc"
)


type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

// NewClient creates a new gRPC client for the Account service
// url: the server address in the format "host:port"
// Returns a pointer to Client and an error if any
// Example usage:
// client, err := account.NewClient("localhost:8080")
func NewClient(url string)(*Client, error){
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to server: " + err.Error())
	}
	
	c := pb.NewAccountServiceClient(conn)

	return &Client{
		conn:    conn,
		service: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}


func (c *Client) PostAccount(ctx context.Context, name, email, password string)(*pb.Account, error){
	req, err := c.service.PostAccount(ctx, &pb.PostAccountRequest{
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return &Account{
		ID:       req.Account.Id,
		Name:     req.Account.Name,
		Email:    req.Account.Email,
	}, nil
}