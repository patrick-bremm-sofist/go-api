package routes

import (
	"go-api/controllers"
	"go-api/serviceContainer"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	userController := serviceContainer.ServiceContainer().InjectUserController()

	app.Post("/user", userController.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Delete("/user/:userId", controllers.DeleteAUser)
}
