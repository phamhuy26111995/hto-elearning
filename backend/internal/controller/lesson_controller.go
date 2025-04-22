package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
)

type LessonController struct {
	lessonService service.LessonService
}

func NewLessonController(lessonService service.LessonService) LessonController {
	return LessonController{
		lessonService: lessonService,
	}
}

func (l LessonController) GetLessonsByModuleId(context *gin.Context) ([]model.Lessons, error) {

	moduleId, _ := context.Params.Get("moduleId")
	moduleIdInt, _ := strconv.ParseInt(moduleId, 10, 64)

	return l.lessonService.GetLessonsByModuleId(moduleIdInt)
}

func (l LessonController) CreateLessons(context *gin.Context) error {

	var lessons []model.Lessons
	if err := context.ShouldBindJSON(&lessons); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	if len(lessons) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Lessons cannot be empty"})
		return nil
	}

	moduleId := lessons[0].ModuleId

	return l.lessonService.CreateLessons(lessons, moduleId)
}

func (l LessonController) UpdateLessons(context *gin.Context) error {

	var lessons []model.Lessons
	if err := context.ShouldBindJSON(&lessons); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	return l.lessonService.UpdateLessons(lessons)
}
