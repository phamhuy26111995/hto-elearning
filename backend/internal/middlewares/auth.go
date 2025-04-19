package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/utils"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, role, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	if role != "TEACHER" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not a teacher."})
		return
	}

	context.Set("userId", userId)
	context.Set("role", role)
	context.Next()
}
