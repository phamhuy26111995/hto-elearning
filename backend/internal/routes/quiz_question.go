package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"

	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func RegisterQuizQuestionController() *controller.QuizQuestionController {
	quizQuestionRepository := repository.NewQuizQuestionRepository()
	quizQuestionService := service.NewQuizQuestionService(quizQuestionRepository)

	return controller.NewQuizQuestionController(quizQuestionService)
}
