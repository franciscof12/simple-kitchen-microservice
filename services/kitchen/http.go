package main

import (
	orders "github.com/franciscof12/kitchen-microservice/protobuf"
	context "golang.org/x/net/context"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (server *httpServer) Run() error {
	router := http.NewServeMux()
	connection := NewGrpcClient(":9000")
	defer connection.Close()
	router.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
		client := orders.NewOrderServiceClient(connection)
		ctx, cancel := context.WithTimeout(request.Context(), time.Second*2)
		defer cancel()
		_, err := client.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 42,
			ProductID:  10,
			Quantity:   3,
		})
		if err != nil {
			slog.Error("failed to create order: %v", err.Error())
			http.Error(write, "failed to create order", http.StatusInternalServerError)
			return
		}
		response, err := client.GetOrder(ctx, &orders.GetOrderRequest{
			OrderID: 42,
		})
		if err != nil {
			slog.Error("failed to get order: %v", err.Error())
			http.Error(write, "failed to get order", http.StatusInternalServerError)
			return
		}
		htmlTemplate := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := htmlTemplate.Execute(write, response.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})
	slog.Info("http server is running on: ", "port", server.addr)
	return http.ListenAndServe(server.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
