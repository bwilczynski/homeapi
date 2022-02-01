package lights

import "context"

type lightsService struct {
	UnimplementedLightServiceServer
}

func NewServer() *lightsService {
	return &lightsService{}
}

func (s *lightsService) List(context.Context, *ListQuery) (*LightList, error) {
	return &LightList{
		Lights: []*Light{
			{Name: "Light 1"},
			{Name: "Light 2"},
			{Name: "Light 3"},
		},
	}, nil
}
