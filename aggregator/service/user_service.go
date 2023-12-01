package service

import (
	"context"
	"time"

	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain/construct"
	"github.com/rezaif79-ri/go-grpc-101/model/user"
)

type userService struct {
	Client user.UserServiceClient
}

func NewUserService(serviceClient user.UserServiceClient) domain.UserService {
	return userService{
		Client: serviceClient,
	}
}

// GreetUser implements domain.UserService.
func (us userService) GreetUser(body construct.GreetUserReq) (construct.GreetUserRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := us.Client.GreetUser(ctx, &user.GreetingRequest{
		Name:  body.Name,
		Salut: body.Salutation,
	})
	if err != nil {
		return construct.GreetUserRes{}, err
	}
	return construct.GreetUserRes{
		Greetings: res.GetGreetingMessage(),
	}, nil
}
