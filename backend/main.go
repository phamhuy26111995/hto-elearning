package main

import (
	_ "github.com/phamhuy26111995/hto-elearning/docs"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/routes"
)

// @title 		E-learning Api
// @version 	1.0
// @description This is E-learning Api
// @termsOfService none

// @contact.name   Huy Pham
// @contact.email  pham.huy.19951126@gmail.com
func main() {
	database.InitDB()

	server := routes.RegisterRoutes()

	err := server.Run(":8080")

	if err != nil {
		return
	}
}
