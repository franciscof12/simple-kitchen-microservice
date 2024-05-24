package handler

import (
	orders "github.com/franciscof12/kitchen-microservice/protobuf"
	"github.com/franciscof12/kitchen-microservice/services/common/util"
	"github.com/franciscof12/kitchen-microservice/services/orders/types"
	"net/http"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewHttpOrderService(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}
	return handler
}
func (handler *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", handler.CreateOrder)
}
func (handler *OrdersHttpHandler) CreateOrder(write http.ResponseWriter, request *http.Request) {
	var requestPayload orders.CreateOrderRequest
	err := util.ParseJSON(request, &requestPayload)
	if err != nil {
		util.WriteError(write, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: requestPayload.GetCustomerID(),
		ProductID:  requestPayload.GetProductID(),
		Quantity:   requestPayload.GetQuantity(),
	}
	err = handler.orderService.CreateOrder(request.Context(), order)
	if err != nil {
		util.WriteError(write, http.StatusInternalServerError, err)
		return
	}
	response := &orders.OrderResponse{Status: "Order created successfully"}
	util.WriteJSON(write, http.StatusCreated, response)
}
