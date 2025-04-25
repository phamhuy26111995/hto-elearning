package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func registerQuizController() *controller.QuizController {

	quizRepo := repository.NewQuizRepository()
	quizService := service.NewQuizService(quizRepo)
	quizController := controller.NewQuizController(quizService)

	return quizController
}
