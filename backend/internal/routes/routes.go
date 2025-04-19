package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	userController := registerUserController()
	courseController := registerCourseController()
	moduleController := registerModuleController()

	authenticated := server.Group("/api/v1/teacher")

	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/users", userController.GetUsers)
	authenticated.POST("/users/create", userController.CreateUser)
	authenticated.POST("/users/create-student", userController.CreateStudent)
	authenticated.PUT("/users/update", userController.UpdateUser)
	authenticated.GET("/users/:id", userController.GetUserById)

	authenticated.GET("/courses", courseController.GetAllCourses)
	authenticated.POST("/course/create", courseController.CreateCourse)
	authenticated.PUT("/course/update", courseController.UpdateCourse)
	authenticated.GET("/course/:id", courseController.GetCourse)
	authenticated.DELETE("/course/:id", courseController.DeleteCourse)

	authenticated.GET("/modules", moduleController.GetAllModulesByCourse)
	authenticated.POST("/modules/create", moduleController.CreateModules)

	server.POST("/login", userController.Login)

	return server
}
