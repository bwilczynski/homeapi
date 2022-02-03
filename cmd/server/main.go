package main

import (
	"flag"
	"os"

	"github.com/bwilczynski/home-api/server"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	hueHost     = os.Getenv("HUE_HOST")
	hueUsername = os.Getenv("HUE_USER")
)

func main() {
	flag.Parse()

	s := server.NewServer(*port, server.HueConfig{Host: hueHost, Username: hueUsername})
	s.Run()
}
