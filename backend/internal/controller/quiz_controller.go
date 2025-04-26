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

func (q *QuizController) GetQuizzesByModuleId(context *gin.Context) {
	moduleIdStr := context.Param("moduleId")

	moduleId, err := strconv.ParseInt(moduleIdStr, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid module ID"})
	}

	quizzes, err := q.quizService.GetAllQuizzesByModuleId(moduleId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quizzes"})
	}

	context.JSON(http.StatusOK, quizzes)
}

func (q *QuizController) CreateQuizzes(context *gin.Context) {
	var quizzes []model.Quiz
	if err := context.ShouldBindJSON(&quizzes); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(quizzes) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Quizzes cannot be empty"})
		return
	}

	err := q.quizService.CreateQuizzes(quizzes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quizzes"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Quizzes created successfully"})
}

func (q *QuizController) UpdateQuizzes(context *gin.Context) {
	var quizzes []model.Quiz
	if err := context.ShouldBindJSON(&quizzes); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(quizzes) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Quizzes cannot be empty"})
		return
	}

	err := q.quizService.UpdateQuizzes(quizzes)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quizzes"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Quizzes updated successfully"})
}
