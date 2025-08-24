package main

import (
	"log"
	"net"

	handler "github.com/itsgitz/grpc-lab/go/services/orders/handler/orders"
	services "github.com/itsgitz/grpc-lab/go/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	//
	// register our gRPC services
	orderService := services.NewOrderService()
	handler.NewGRPCOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC server on port", s.addr)

	return grpcServer.Serve(lis)
}
