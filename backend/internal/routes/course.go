package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func registerCourseController() *controller.CourseController {
	courseRepo := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepo)
	courseController := controller.NewCourseController(courseService)

	return &courseController
}
