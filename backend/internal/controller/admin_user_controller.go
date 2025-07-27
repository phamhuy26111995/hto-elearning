package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/constant"
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
)

type AdminUserController struct {
	service service.AdminUserService
}

func NewAdminUserController(service service.AdminUserService) *AdminUserController {
	return &AdminUserController{
		service: service,
	}
}

// CreateUser godoc
// @Summary Create a new user (teacher or student)
// @Description Create a new user with role TEACHER or STUDENT. Only accessible by ADMIN.
// @Tags Admin Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body dto.UserDTO true "User info to create"
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]string "Invalid input or role"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /admin/user/create [post]
func (controller *AdminUserController) CreateUser(context *gin.Context) {
	userId, exists := context.Get("userId")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var userDto dto.UserDTO
	if err := context.ShouldBindJSON(&userDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDto.CreatedBy = userId.(int64)
	userDto.UpdatedBy = userId.(int64)

	if userDto.Status == "" {
		userDto.Status = "ACTIVE"
	}

	if userDto.Role != constant.RoleTeacher && userDto.Role != constant.RoleStudent {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	userEntity := userDto.MappingToUserEntity(true)

	if err := controller.service.CreateUser(userEntity); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"User created successfully": userEntity.UserID})

}

func (controller *AdminUserController) GetAllTeachers(context *gin.Context) {
	var paging dto.Paging

	if err := context.ShouldBindQuery(&paging); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teachers, err := controller.service.GetAllByRole(constant.RoleTeacher, paging)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": teachers})
}

func (controller *AdminUserController) GetAllStudents(context *gin.Context) {
	var paging dto.Paging

	if err := context.ShouldBindQuery(&paging); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teachers, err := controller.service.GetAllByRole(constant.RoleStudent, paging)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": teachers})
}
