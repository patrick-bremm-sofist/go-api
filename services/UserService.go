package services

import (
	"go-api/interfaces"
	"go-api/models"
	"go-api/responses"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServices struct {
	interfaces.IUserRepository
}

var validate = validator.New()

func (service *UserServices) CreateAUser(c *fiber.Ctx) error {
	var user models.User

	// Validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	// Use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	result, err := service.InsertOne(newUser)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(
		responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"data": result}})
}
