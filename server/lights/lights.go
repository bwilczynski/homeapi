package lights

import (
	"context"

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

func (s *lightsService) List(context.Context, *ListQuery) (*LightList, error) {
	lights, err := s.bridge.GetLights()
	if err != nil {
		return nil, err
	}
	res := make([]*Light, len(lights))
	for i, light := range lights {
		res[i] = &Light{Name: light.Name}
	}

	return &LightList{
		Lights: res,
	}, nil
}
