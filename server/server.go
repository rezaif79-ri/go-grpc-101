package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rezaif79-ri/go-grpc-101/model/user"
	"google.golang.org/grpc"
)

// userService is a struct implementing UserServiceServer
type userService struct{}

// GreetUser return greeting message given the name and salutation
// in gRPC protocol
func (*userService) GreetUser(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {
	//business logic
	salutationMessage := fmt.Sprintf("Howdy, %s %s, nice to see you in the future!",
		req.Salut, req.Name)

	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	server := grpc.NewServer()

	user.RegisterUserServiceServer(server, &userService{})

	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
