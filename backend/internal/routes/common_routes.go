package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func CommonRoutes(server *gin.Engine) {

	userController := RegisterUserController()

	apiGroup := server.Group("/api/v1/common")

	apiGroup.Use(middlewares.Authenticate)

	apiGroup.GET("/current-user", userController.GetCurrentUserLogin)

}
