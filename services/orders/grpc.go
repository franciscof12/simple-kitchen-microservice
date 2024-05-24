package main

import (
	handler "github.com/franciscof12/kitchen-microservice/services/orders/handler/orders"
	"github.com/franciscof12/kitchen-microservice/services/orders/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type GrpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *GrpcServer {
	return &GrpcServer{addr: addr}
}

func (server *GrpcServer) Run() error {
	listener, err := net.Listen("tcp", server.addr)
	if err != nil {
		slog.Error("failed to listen: %v", err.Error())
	}
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcServer, orderService)
	slog.Info("grpc server is running on %s", server.addr)
	return grpcServer.Serve(listener)
}
