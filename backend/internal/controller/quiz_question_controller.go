package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
)

type QuizQuestionController struct {
	quizQuestionService service.QuizQuestionService
}

func NewQuizQuestionController(quizQuestionService service.QuizQuestionService) *QuizQuestionController {
	return &QuizQuestionController{
		quizQuestionService: quizQuestionService,
	}
}

func (c *QuizQuestionController) GetAllQuestionsByQuizId(context *gin.Context) {
	quizIdStr := context.Param("quizId")
	quizId, err := strconv.ParseInt(quizIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	questions := c.quizQuestionService.GetAllQuestionsByQuizId(quizId)

	context.JSON(http.StatusOK, questions)
}

func (c *QuizQuestionController) CreateQuestions(context *gin.Context) {
	var questions []model.QuizQuestion
	err := context.ShouldBindJSON(&questions)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(questions) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Questions cannot be empty"})
		return
	}

	quizId := questions[0].QuizId

	err = c.quizQuestionService.CreateQuestions(quizId, questions)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Questions created successfully"})
}

func (c *QuizQuestionController) UpdateQuestions(context *gin.Context) {
	var questions []model.QuizQuestion
	err := context.ShouldBindJSON(&questions)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(questions) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Questions cannot be empty"})
		return
	}

	err = c.quizQuestionService.UpdateQuestions(questions)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Questions updated successfully"})
}
