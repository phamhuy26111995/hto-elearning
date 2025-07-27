package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// ShowUsers godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.User
// @Router       /users [get]
func (controller *UserController) GetUsers(context *gin.Context) {
	users, err := controller.userService.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func (controller *UserController) GetUsersByTeacher(context *gin.Context) {
	userId, exists := context.Get("userId")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	users, err := controller.userService.GetAllUsersByTeacherId(userId.(int64))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func (controller *UserController) CreateUser(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.userService.CreateUser(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "User created successfully"})
}

func (controller *UserController) CreateStudent(context *gin.Context) {
	userId, exists := context.Get("userId")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.userService.CreateStudent(&user, userId.(int64)); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Student created successfully"})
}

func (controller *UserController) GetCurrentUserLogin(context *gin.Context) {
	userIdFromContext, _ := context.Get("userId")

	userId := userIdFromContext.(int64)

	user, err := controller.userService.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"userInfo": user})
}

func (controller *UserController) GetUserById(context *gin.Context) {
	userIdParam := context.Param("id")

	userId, parseErr := strconv.ParseInt(userIdParam, 10, 64)

	if parseErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": parseErr.Error()})
		return
	}

	user, err := controller.userService.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func (controller *UserController) UpdateUser(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.userService.UpdateUser(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func (controller *UserController) UpdateStudent(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rawUserId, exists := context.Get("userId")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user.UpdatedBy = rawUserId.(int64)

	if err := controller.userService.UpdateUser(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user_id": user.UserID})
}

func (controller *UserController) Login(context *gin.Context) {
	var user dto.UserLoginDTO
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, userInfo := controller.userService.Login(&user)

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Invalid username or password"})
		return
	}

	userInfoDto := dto.UserDTO{UserID: userInfo.UserID, Username: userInfo.Username, Email: userInfo.Email, Role: userInfo.Role,
		CreatedAt: userInfo.CreatedAt, CreatedBy: userInfo.CreatedBy, UpdatedAt: userInfo.UpdatedAt, UpdatedBy: userInfo.UpdatedBy}

	context.JSON(http.StatusOK, gin.H{"token": token, "userInfo": userInfoDto})
}

func (controller *UserController) EnrollCourseForStudent(context *gin.Context) {
	var enrollment dto.EnrollmentDTO
	if err := context.ShouldBindJSON(&enrollment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.userService.EnrollCourseForStudent(enrollment.UserId, enrollment.CourseId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Enrollment": "Enrolled successfully"})
}

func (controller *UserController) UnEnrollCourseForStudent(context *gin.Context) {
	var enrollment dto.EnrollmentDTO
	if err := context.ShouldBindJSON(&enrollment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.userService.UnEnrollCourseForStudent(enrollment.UserId, enrollment.CourseId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Enrollment": "UnEnrolled successfully"})
}

func (controller *UserController) ChangeStatus(context *gin.Context) {
	var student dto.UserDTO
	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.userService.ChangeStatus(student.UserID, student.Status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Success": "Status changed successfully"})
}

func (controller *UserController) DeleteUser(context *gin.Context) {
	param := context.Param("id")

	rawUserId, exists := context.Get("userId")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorize"})
		return
	}

	userId, parseErr := strconv.ParseInt(param, 10, 64)

	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Parse Param Failed"})
		return
	}

	var teacherId int64

	teacherId = rawUserId.(int64)

	err := controller.userService.DeleteUserPermanently(userId, teacherId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Success": "Delete successfully"})
}
