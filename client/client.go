package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rezaif79-ri/go-grpc-101/model/user"
	"google.golang.org/grpc"
)

const userServiceAddress = "localhost:5001"

func main() {
	//Create connection to gRPC server
	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	//Create new userService client
	userServiceClient := user.NewUserServiceClient(conn)

	//create connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := userServiceClient.GreetUser(ctx, &user.GreetingRequest{
		Name:  "Marty Mc Fly",
		Salut: "Mr.",
	})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	//show response
	fmt.Println(response.GreetingMessage)
}
