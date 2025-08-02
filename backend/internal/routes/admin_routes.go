package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func AdminRoutes(server *gin.Engine) {

	userController := RegisterAdminUserController()

	courseController := RegisterCourseController()

	//moduleController := RegisterModuleController()

	apiGroup := server.Group("/api/v1/admin")

	apiGroup.Use(middlewares.Authenticate, middlewares.AuthorizeAdmin)

	userApiGroup := apiGroup.Group("/user")

	userApiGroup.GET("/teacher", userController.GetAllTeachers)

	userApiGroup.GET("/student", userController.GetAllStudents)

	userApiGroup.POST("/create", userController.CreateUser)

	userApiGroup.PUT("/update", userController.UpdateUser)

	courseApiGroup := apiGroup.Group("/course")

	courseApiGroup.GET("/", userController.GetAllTeachers)

	courseApiGroup.POST("/create", courseController.CreateCourse)

	courseApiGroup.PUT("/update", courseController.UpdateCourse)

}
