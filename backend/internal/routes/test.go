package routes

import "github.com/gin-gonic/gin"

func getTest(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}
