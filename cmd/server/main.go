package main

import (
	"flag"
	"os"

	"github.com/bwilczynski/homeapi"
	"github.com/bwilczynski/homeapi/lights"
	"github.com/bwilczynski/homeapi/lights/hue"
	"google.golang.org/grpc"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	hueHost     = os.Getenv("HUE_HOST")
	hueUsername = os.Getenv("HUE_USER")
)

func main() {
	flag.Parse()

	s := homeapi.NewServer(*port)
	s.Run(func(s *grpc.Server) {
		lights.RegisterLightServiceServer(s, hue.NewServer(hueHost, hueUsername))
	})
}
