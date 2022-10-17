package routes

import (
	"go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Delete("/user/:userId", controllers.DeleteAUser)
}
