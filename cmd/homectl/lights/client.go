package lights

import (
	"context"
	"fmt"

	"github.com/bwilczynski/homeapi/lights"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type lightServiceClient struct {
	target *string
}

func (c *lightServiceClient) dial() (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(*c.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err = fmt.Errorf("error while dialing %w", err)
	}
	return
}

func (c *lightServiceClient) ToggleGroup(ctx context.Context, groupId string) (err error) {
	conn, err := c.dial()
	if err != nil {
		return
	}
	defer conn.Close()

	client := lights.NewLightServiceClient(conn)
	_, err = client.ToggleGroup(ctx, &lights.ToggleGroupRequest{GroupId: groupId})
	return
}

func (c *lightServiceClient) List(ctx context.Context) (ll []*lights.Light, err error) {
	conn, err := c.dial()
	if err != nil {
		return
	}
	defer conn.Close()

	client := lights.NewLightServiceClient(conn)
	lr, err := client.List(ctx, &lights.ListRequest{})
	if err != nil {
		return
	}

	return lr.Lights, nil
}
