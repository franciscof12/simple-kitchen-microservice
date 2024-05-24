package main

import (
	handler "github.com/franciscof12/kitchen-microservice/services/orders/handler/orders"
	"github.com/franciscof12/kitchen-microservice/services/orders/service"
	"log/slog"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (server *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderService(orderService)
	orderHandler.RegisterRouter(router)
	slog.Info("http server is running on: ", "port", server.addr)
	return http.ListenAndServe(server.addr, router)
}
