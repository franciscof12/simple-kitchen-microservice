package handler

import (
	"context"
	orders "github.com/franciscof12/kitchen-microservice/protobuf"
	"github.com/franciscof12/kitchen-microservice/services/orders/types"
	"google.golang.org/grpc"
	"log/slog"
)

type OrderGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpcServer *grpc.Server, orderService types.OrderService) {
	grpcHandler := &OrderGrpcHandler{
		orderService: orderService,
	}
	orders.RegisterOrderServiceServer(grpcServer, grpcHandler)
}

func (handler *OrderGrpcHandler) GetOrder(ctx context.Context, request *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	// ors mean orders
	ors := handler.orderService.GetOrder(ctx)
	response := &orders.GetOrderResponse{
		Orders: ors,
	}
	return response, nil
}

func (handler *OrderGrpcHandler) CreateOrder(context context.Context, request *orders.CreateOrderRequest) (*orders.OrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   4,
	}
	if err := handler.orderService.CreateOrder(context, order); err != nil {
		slog.Error("failed to create order: %v", err.Error())
	}
	response := &orders.OrderResponse{
		Status: "Order created successfully",
	}
	return response, nil
}
