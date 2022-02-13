package hue

import (
	"context"
	"strconv"

	"github.com/amimof/huego"
	l "github.com/bwilczynski/homeapi/lights"
)

type lightsService struct {
	bridge *huego.Bridge

	l.UnimplementedLightServiceServer
}

func NewServer(host, username string) *lightsService {
	server := &lightsService{
		bridge: huego.New(host, username),
	}

	return server
}

func (s *lightsService) List(ctx context.Context, req *l.ListRequest) (*l.ListResponse, error) {
	lights, err := s.bridge.GetLightsContext(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*l.Light, len(lights))
	for i, light := range lights {
		res[i] = &l.Light{Id: strconv.Itoa(light.ID), Name: light.Name}
	}

	return &l.ListResponse{
		Lights: res,
	}, nil
}

func (s *lightsService) ListGroups(ctx context.Context, req *l.ListGroupsRequest) (*l.ListGroupsResponse, error) {
	groups, err := s.bridge.GetGroupsContext(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*l.Group, len(groups))
	for i, group := range groups {
		res[i] = &l.Group{Id: strconv.Itoa(group.ID), Name: group.Name, Lights: group.Lights}
	}

	return &l.ListGroupsResponse{
		Groups: res,
	}, nil
}

func (s *lightsService) ToggleGroup(ctx context.Context, req *l.ToggleGroupRequest) (*l.ToggleGroupResponse, error) {
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

	return &l.ToggleGroupResponse{}, nil
}
