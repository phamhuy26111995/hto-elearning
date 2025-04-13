package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func registerUserController() *controller.UserController {

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return userController
}
