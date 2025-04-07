package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	public := server.Group("/")

	public.GET("", func(c *gin.Context) {})

	authenticated := server.Group("/api")

	registerUserRoutes(authenticated)

}
