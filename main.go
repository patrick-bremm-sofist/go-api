package main

import (
	"go-api/configs"
	"go-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Run database
	configs.ConnectDB()

	// Routes
	routes.UserRoute(app)

	app.Listen(":6000")
}
