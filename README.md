# GO GRPC 101

Build gRPC for service app and REST-API for aggregator app
Trying to implement clean architecture on aggregator

## Summary:

- Install protobuf-compiler (libprotoc)
- Run `go get -u github.com/golang/protobuf/protoc-gen-go`
- Run `go get -u google.golang.org/golang/protobuf`
- Create your proto definition in proto folders
- Run `protoc --go_out=. .\proto\user.proto` & `protoc --go-grpc_out=. .\proto\user.proto`
- Run `protoc --go_out=. .\proto\calculator.proto` & `protoc --go-grpc_out=. .\proto\calculator.proto`
- Run `go run server/server.go` for gRPC service & `go run aggregator/server.go` for rest-api

## How to run this?

- Make sure you have docker installed
- Run `docker-compose up -d`
- Test the api on localhost:6001