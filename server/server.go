package main

import (
	"log"
	"net"

	"github.com/rezaif79-ri/go-grpc-101/model/calculator"
	"github.com/rezaif79-ri/go-grpc-101/model/user"
	"github.com/rezaif79-ri/go-grpc-101/server/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	server := grpc.NewServer()
	calculatorService := service.NewCalculatorServices()
	userService := service.NewUserService()

	user.RegisterUserServiceServer(server, userService)
	calculator.RegisterCalculatorServiceServer(server, calculatorService)
	log.Println("SERVER: running at tcp :5001")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
