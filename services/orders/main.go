package main

import "log/slog"

const (
	addr = ":9000"
)

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := NewGRPCServer(addr)
	err := grpcServer.Run()
	if err != nil {
		slog.Error("failed to run grpc server: %v", err.Error())
	}
}
