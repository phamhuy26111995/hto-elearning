package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
)

type CourseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return CourseController{courseService: courseService}
}

func (c *CourseController) GetAllCourses(ctx *gin.Context) {
	raw, _ := ctx.Get("userId")
	userId, ok := raw.(int64)
	if !ok {
		ctx.JSON(500, "Error when get context value userId")
		return
	}
	courses, err := c.courseService.GetAllCourses(userId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, courses)
}

func (c *CourseController) GetCourse(ctx *gin.Context) {
	courseId := ctx.Param("id")
	id, parseErr := strconv.ParseInt(courseId, 10, 64)
	if parseErr != nil {
		ctx.JSON(400, parseErr)
		return
	}
	course, err := c.courseService.GetCourse(id)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, course)
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBindJSON(&course)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	err = c.courseService.CreateCourse(&course, course.TeacherId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, gin.H{"Success": "Course created successfully"})
}

func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	var course model.Course
	err := ctx.ShouldBindJSON(&course)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	err = c.courseService.UpdateCourse(&course)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, gin.H{"Success": "Course updated successfully"})
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	courseId := ctx.Param("id")
	id, parseErr := strconv.ParseInt(courseId, 10, 64)
	if parseErr != nil {
		ctx.JSON(400, parseErr)
		return
	}
	err := c.courseService.DeleteCourse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
