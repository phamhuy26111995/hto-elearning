package main

import (
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/routes"
)

func main() {
	database.InitDB()

	server := routes.RegisterRoutes()

	err := server.Run(":8080")

	if err != nil {
		return
	}
}
