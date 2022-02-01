package main

import (
	"log"

	"github.com/bwilczynski/home-api/server"
)

func main() {
	log.Printf("Starting lights server.")

	s := server.NewServer()
	s.Run()
}
