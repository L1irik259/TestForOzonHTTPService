package client

import (
	"context"
	"time"

	pb "github.com/L1irik259/TestForOzonHTTPService/internal/transport/proto/github.com/L1irik259/TestForOzon/transport/genetation/go/v1"
	"google.golang.org/grpc"
)

type ItemServiceClient struct {
	client pb.OzonServiceClient
}

func NewItemServiceClient(conn *grpc.ClientConn) *ItemServiceClient {
	return &ItemServiceClient{
		client: pb.NewOzonServiceClient(conn),
	}
}

func (c *ItemServiceClient) FindAllItemsByDate(date time.Time) ([]*pb.Item, error) {
	req := &pb.ItemRequest{
		Date: date.Format("02/01/2006"),
	}

	resp, err := c.client.GetItem(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
