package server

import (
	"fmt"
	"log"
	"net"

	"github.com/bwilczynski/home-api/server/lights"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type homeApiServer struct {
}

func NewServer() *homeApiServer {
	return &homeApiServer{}
}

func (s *homeApiServer) Run(port int) {
	addr := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Started GRPC server on %v", addr)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	lights.RegisterLightServiceServer(grpcServer, lights.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
