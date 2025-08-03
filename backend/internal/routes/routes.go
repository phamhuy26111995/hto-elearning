package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userController := RegisterUserController()

	CommonRoutes(server)

	TeacherRoutes(server)

	AdminRoutes(server)

	server.POST("/login", userController.Login)

	//server.GET("/api/v1/common/current-user", userController.GetCurrentUserLogin)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return server
}
