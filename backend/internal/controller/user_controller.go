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

func (controller *UserController) GetUserById(context *gin.Context) {
	userId := context.Param("id")
	parseInt, parseErr := strconv.ParseInt(userId, 10, 64)
	if parseErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	user, err := controller.userService.GetUserById(parseInt)

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

func (controller *UserController) Login(context *gin.Context) {
	var user dto.UserLoginDTO
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := controller.userService.Login(&user)

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Invalid username or password"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
