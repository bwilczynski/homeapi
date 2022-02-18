package homeapi

import (
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type homeApiServer struct {
	port int
}

type ServiceRegistrar func(*grpc.Server)

func NewServer(port int) *homeApiServer {
	return &homeApiServer{
		port: port,
	}
}

func (s *homeApiServer) Run(rr ...ServiceRegistrar) {
	addr := fmt.Sprintf("localhost:%d", s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Started GRPC server on %v", addr)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	for _, r := range rr {
		r(grpcServer)
	}

	for k, v := range grpcServer.GetServiceInfo() {
		mm := make([]string, len(v.Methods))
		for i, m := range v.Methods {
			mm[i] = m.Name
		}
		log.Printf("Service: %s (%s)", k, strings.Join(mm, ", "))
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
