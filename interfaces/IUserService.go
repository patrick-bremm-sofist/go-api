package interfaces

import "github.com/gofiber/fiber/v2"

type IUserService interface {
	CreateAUser(c *fiber.Ctx) error
}
