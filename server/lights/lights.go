package lights

import (
	"context"
	"strconv"

	"github.com/amimof/huego"
)

type lightsService struct {
	bridge *huego.Bridge

	UnimplementedLightServiceServer
}

func NewServer(host, username string) *lightsService {
	server := &lightsService{
		bridge: huego.New(host, username),
	}

	return server
}

func (s *lightsService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	lights, err := s.bridge.GetLightsContext(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*Light, len(lights))
	for i, light := range lights {
		res[i] = &Light{Id: strconv.Itoa(light.ID), Name: light.Name}
	}

	return &ListResponse{
		Lights: res,
	}, nil
}

func (s *lightsService) ListGroups(ctx context.Context, req *ListGroupsRequest) (*ListGroupsResponse, error) {
	groups, err := s.bridge.GetGroupsContext(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*Group, len(groups))
	for i, group := range groups {
		res[i] = &Group{Id: strconv.Itoa(group.ID), Name: group.Name, Lights: group.Lights}
	}

	return &ListGroupsResponse{
		Groups: res,
	}, nil
}

func (s *lightsService) ToggleGroup(ctx context.Context, req *ToggleGroupRequest) (*ToggleGroupResponse, error) {
	id, err := strconv.Atoi(req.GroupId)
	if err != nil {
		return nil, err
	}

	group, err := s.bridge.GetGroupContext(ctx, id)
	if err != nil {
		return nil, err
	}

	if group.IsOn() {
		group.OffContext(ctx)
	} else {
		group.OnContext(ctx)
	}

	return &ToggleGroupResponse{}, nil
}
