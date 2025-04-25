package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
)

type QuizController struct {
	quizService service.QuizService
}

func NewQuizController(quizService service.QuizService) *QuizController {
	return &QuizController{
		quizService: quizService,
	}
}

func (q *QuizController) GetQuizzesByModuleId(context *gin.Context) ([]model.Quiz, error) {
	moduleIdStr := context.Param("moduleId")

	moduleId, err := strconv.ParseInt(moduleIdStr, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid module ID"})
	}

	return q.quizService.GetAllQuizzesByModuleId(moduleId)
}
