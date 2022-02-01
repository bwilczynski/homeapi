package main

import (
	"flag"

	"github.com/bwilczynski/home-api/server"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.Run(*port)
}
