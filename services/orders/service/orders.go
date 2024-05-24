package service

import (
	order "github.com/franciscof12/kitchen-microservice/protobuf"
	"golang.org/x/net/context"
)

var ordersDB = make([]*order.Order, 0)

type OrderService struct {
	// order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (service *OrderService) CreateOrder(context context.Context, order *order.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}

func (service *OrderService) GetOrder(context context.Context) []*order.Order {
	return ordersDB
}
