run-orders:
	@go run services/orders/*.go

run-kitchen:
	@go run services/kitchen/*.go

gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protobuf/orders.proto