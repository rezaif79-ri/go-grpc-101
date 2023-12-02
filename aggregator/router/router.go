package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/config"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/controller"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/service"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/util"
	"github.com/rezaif79-ri/go-grpc-101/model/calculator"
	"github.com/rezaif79-ri/go-grpc-101/model/user"
)

func SetupRouter(app *fiber.App) {
	// Init utilities
	validator := util.NewGoValidator()

	// Init tcp conn
	conn, err := config.NewServiceGrpcConn()
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		conn.Close()
		return
	}
	// Init model
	var userModel = user.NewUserServiceClient(conn)
	var calcModel = calculator.NewCalculatorServiceClient(conn)

	// Init service
	var userService = service.NewUserService(userModel)
	var calcService = service.NewCalculatorService(calcModel)

	// Init controller
	var userController = controller.NewUserController(&userService)
	var calcController = controller.NewCalculatorController(*validator, &calcService)

	// Set route and group
	var userRoute = app.Group("/api/users")
	userRoute.Post("greetings", userController.GreetUser)

	var calculatorRoute = app.Group("/api/calculators")
	calculatorRoute.Post("", calcController.Calculate)
}
