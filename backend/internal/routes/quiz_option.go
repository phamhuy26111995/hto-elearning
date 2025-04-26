package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func registerQuizOptionController() *controller.QuizOptionController {
	quizOptionRepo := repository.NewQuizOptionRepository()
	quizOptionService := service.NewQuizOptionService(quizOptionRepo)

	return controller.NewQuizOptionController(quizOptionService)
}
