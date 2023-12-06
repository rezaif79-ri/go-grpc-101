package service

import (
	"context"
	"fmt"

	"github.com/rezaif79-ri/go-grpc-101/server/model/user"
)

// userService is a struct implementing UserServiceServer
type userService struct {
	user.UnimplementedUserServiceServer
}

func NewUserService() user.UserServiceServer {
	return &userService{
		UnimplementedUserServiceServer: user.UnimplementedUserServiceServer{},
	}
}

// GreetUser return greeting message given the name and salutation
// in gRPC protocol
func (us *userService) GreetUser(ctx context.Context, req *user.GreetingRequest) (*user.GreetingResponse, error) {
	//business logic
	salutationMessage := fmt.Sprintf("Howdy, %s %s, nice to see you in the future!",
		req.Salut, req.Name)

	return &user.GreetingResponse{GreetingMessage: salutationMessage}, nil
}
