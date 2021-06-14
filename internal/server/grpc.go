package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	*grpc.Server
	port uint
}

func NewGrpcServer(port uint) *GrpcServer {
	grpcServer := grpc.NewServer()

	return &GrpcServer{
		Server: grpcServer,
		port:   port,
	}
}

func (s *GrpcServer) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	return s.Server.Serve(listener)
}
