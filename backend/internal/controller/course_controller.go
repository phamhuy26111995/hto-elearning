package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/constant"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
	"time"
)

type CourseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return CourseController{courseService: courseService}
}

func (c *CourseController) GetAll(ctx *gin.Context) {
	raw, exists := ctx.Get("role")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "role not exist"})
	}

	role := raw.(string)

	if role != constant.RoleAdmin {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid role"})
	}

	courses, err := c.courseService.GetAll()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, courses)
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
	userId, exist := ctx.Get("userId")

	if !exist {
		ctx.JSON(400, "Error when get context value userId")
		return
	}

	var course model.Course
	err := ctx.ShouldBindJSON(&course)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	course.CreatedBy = userId.(int64)
	course.UpdatedBy = userId.(int64)
	course.UserId = userId.(int64)

	err = c.courseService.CreateCourse(&course, course.UserId)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, gin.H{"Success": "Course created successfully"})
}

func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")

	if !exist {
		ctx.JSON(400, "Error when get context value userId")
		return
	}

	var course model.Course
	err := ctx.ShouldBindJSON(&course)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	course.UpdatedBy = userId.(int64)
	course.UpdatedAt = time.Now()

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
