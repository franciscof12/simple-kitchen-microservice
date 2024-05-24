package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

func NewGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	httpServer := NewHttpServer(":1000")
	err := httpServer.Run()
	if err != nil {
		slog.Error("failed to run http server: %v", "error", err.Error())
	}
}
