package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/domain/construct"
)

type UserController interface {
	GreetUser(c *fiber.Ctx)
}

type UserService interface {
	GreetUser(body construct.GreetUserReq) (construct.GreetUserRes, error)
}
