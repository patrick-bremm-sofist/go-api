package repositories

import (
	"context"
	"go-api/configs"
	"go-api/interfaces"
	"go-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserRepository interfaces.IUserRepository
}

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func (*UserRepository) InsertOne(newUser models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, newUser)

	return result.InsertedID.(primitive.ObjectID).String(), err
}
