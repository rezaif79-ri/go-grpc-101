package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/router"
)

const userServiceAddress = "localhost:5001"

func main() {

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	router.SetupRouter(app)

	app.Listen("127.0.0.1:3210")
}
