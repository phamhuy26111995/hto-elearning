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

func (l *LessonController) GetLessonsByModuleId(context *gin.Context) {

	moduleId := context.Param("moduleId")
	moduleIdInt, _ := strconv.ParseInt(moduleId, 10, 64)

	lessons, err := l.lessonService.GetLessonsByModuleId(moduleIdInt)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, lessons)
}

func (l LessonController) CreateLessons(context *gin.Context) {

	var lessons []model.Lessons
	if err := context.ShouldBindJSON(&lessons); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(lessons) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Lessons cannot be empty"})
		return
	}

	moduleId := lessons[0].ModuleId

	err := l.lessonService.CreateLessons(lessons, moduleId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Lessons created successfully"})

}

func (l LessonController) UpdateLessons(context *gin.Context) {

	var lessons []model.Lessons
	if err := context.ShouldBindJSON(&lessons); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(lessons) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Lessons cannot be empty"})
		return
	}

	err := l.lessonService.UpdateLessons(lessons)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Lessons updated successfully"})

}
