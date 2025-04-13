package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	userController := registerUserController()

	server.GET("/users", userController.GetUsers)
	server.POST("/users/create", userController.CreateUser)
	server.PUT("/users/update", userController.UpdateUser)
	server.GET("/users/:id", userController.GetUserById)

	return server
}
