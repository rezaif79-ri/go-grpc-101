package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/config"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/controller"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/service"
	"github.com/rezaif79-ri/go-grpc-101/model/user"
)

func SetupRouter(app *fiber.App) {
	// Init tcp conn
	conn, err := config.NewServiceGrpcConn()
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		conn.Close()
		return
	}
	// Init model
	var userModel = user.NewUserServiceClient(conn)

	// Init service
	var userService = service.NewUserService(userModel)

	// Init controller
	var userController = controller.NewUserController(&userService)
	// Set route and group
	var user = app.Group("/api/users")
	user.Post("greetings", userController.GreetUser)
}
