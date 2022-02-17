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

func (c *lightServiceClient) ToggleGroup(ctx context.Context, groupId string) (err error) {
	conn, err := grpc.Dial(*c.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("error while dialing %w", err)
	}
	defer conn.Close()

	client := lights.NewLightServiceClient(conn)
	_, err = client.ToggleGroup(ctx, &lights.ToggleGroupRequest{GroupId: groupId})
	return
}
