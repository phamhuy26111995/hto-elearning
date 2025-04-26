package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func registerLessonController() *controller.LessonController {

	lessonRepo := repository.NewLessonRepository()
	lessonService := service.NewLessonService(lessonRepo)
	lessonController := controller.NewLessonController(lessonService)

	return &lessonController
}
