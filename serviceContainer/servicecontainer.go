package serviceContainer

import (
	"go-api/controllers"
	"go-api/repositories"
	"go-api/services"
	"sync"
)

type IServiceContainer interface {
	InjectUserController() controllers.UserController
}

type kernel struct{}

func (k *kernel) InjectUserController() controllers.UserController {
	userRepository := &repositories.UserRepository{}
	userService := &services.UserServices{IUserRepository: userRepository}
	userController := controllers.UserController{IUserService: userService}

	return userController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
