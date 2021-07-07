generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/helloworld.proto

run-server:
	go run server.go

run-client:
	go run client.go
