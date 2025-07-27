package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func AdminRoutes(server *gin.Engine) {

	userController := RegisterAdminUserController()

	apiGroup := server.Group("/api/v1/admin")

	apiGroup.Use(middlewares.Authenticate, middlewares.AuthorizeAdmin)

	userApiGroup := apiGroup.Group("/user")

	userApiGroup.GET("/teacher", userController.GetAllTeachers)

	userApiGroup.GET("/student", userController.GetAllStudents)

	userApiGroup.POST("/create", userController.CreateUser)

	userApiGroup.PUT("/update", userController.UpdateUser)

}
