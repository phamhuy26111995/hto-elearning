package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
)

type QuizOptionController struct {
	quizOptionService service.QuizOptionService
}

func NewQuizOptionController(quizOptionService service.QuizOptionService) *QuizOptionController {
	return &QuizOptionController{
		quizOptionService: quizOptionService,
	}
}

func (c *QuizOptionController) CreateQuizOptions(context *gin.Context) {
	var quizOptions []model.QuizOption
	if err := context.ShouldBindJSON(&quizOptions); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if len(quizOptions) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Quiz options cannot be empty"})
		return
	}

	questionId := quizOptions[0].QuestionId

	err := c.quizOptionService.CreateQuizOptions(quizOptions, questionId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Quiz option created successfully"})
}

func (c *QuizOptionController) UpdateQuizOptions(context *gin.Context) {
	var quizOptions []model.QuizOption
	if err := context.ShouldBindJSON(&quizOptions); err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if len(quizOptions) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Quiz options cannot be empty"})
		return
	}

	err := c.quizOptionService.UpdateQuizOptions(quizOptions)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Quiz option updated successfully"})
}

func (c *QuizOptionController) GetAllQuizOptionsByQuestionId(context *gin.Context) {

	questionIdStr := context.Param("questionId")
	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	quizOptions, err := c.quizOptionService.GetAllQuizOptionsByQuestionId(questionId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, quizOptions)
}
