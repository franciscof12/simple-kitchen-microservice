package types

import (
	"context"
	orders "github.com/franciscof12/kitchen-microservice/protobuf"
)

type OrderService interface {
	CreateOrder(context context.Context, payload *orders.Order) error
	GetOrder(context context.Context) []*orders.Order
}
