package main

import (
	"github.com/gofiber/fiber/v2/middleware/pprof"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rezaif79-ri/go-grpc-101/aggregator/router"
)

func main() {

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(pprof.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	router.SetupRouter(app)

	app.Listen(":6001")
}
