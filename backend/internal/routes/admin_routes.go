package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func AdminRoutes(server *gin.Engine) {

	userController := RegisterUserController()

	adminUserController := RegisterAdminUserController()

	courseController := RegisterCourseController()

	moduleController := RegisterModuleController()

	apiGroup := server.Group("/api/v1/admin")

	apiGroup.Use(middlewares.Authenticate, middlewares.AuthorizeAdmin)

	userApiGroup := apiGroup.Group("/user")

	userApiGroup.GET("/teacher", adminUserController.GetAllTeachers)

	userApiGroup.GET("/student", adminUserController.GetAllStudents)

	userApiGroup.GET("/:id", userController.GetUserById)

	userApiGroup.POST("/create", adminUserController.CreateUser)
	userApiGroup.POST("/creates", adminUserController.CreateUsers)

	userApiGroup.PUT("/update", adminUserController.UpdateUser)

	courseApiGroup := apiGroup.Group("/course")

	courseApiGroup.GET("/", courseController.GetAll)

	courseApiGroup.POST("/create", courseController.CreateCourse)

	courseApiGroup.PUT("/update", courseController.UpdateCourse)

	moduleApiGroup := apiGroup.Group("/module")

	moduleApiGroup.GET("/:courseId", moduleController.GetAllModulesByCourse)

}
