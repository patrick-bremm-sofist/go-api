package interfaces

import (
	"go-api/models"
)

type IUserRepository interface {
	InsertOne(newUser models.User) (string, error)
}
