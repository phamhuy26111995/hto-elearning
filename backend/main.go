package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/routes"
)

func main() {
	database.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
