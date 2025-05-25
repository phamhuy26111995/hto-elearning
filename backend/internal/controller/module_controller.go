package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
	"net/http"
	"strconv"
	"time"
)

type ModuleController struct {
	moduleService     service.ModuleService
	bulkModuleService service.BulkSaveService
}

func NewModuleController(moduleService service.ModuleService, bulkModuleService service.BulkSaveService) ModuleController {
	return ModuleController{moduleService: moduleService, bulkModuleService: bulkModuleService}
}

func (controller *ModuleController) GetAllModulesByCourse(context *gin.Context) {
	courseId := context.Param("courseId")
	parseInt, parseErr := strconv.ParseInt(courseId, 10, 64)

	if parseErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
	}

	modules, err := controller.moduleService.GetAllModulesByCourse(parseInt)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"modules": modules})
}

func (controller *ModuleController) CreateModules(context *gin.Context) {
	var modules []*model.Module
	if err := context.ShouldBindJSON(&modules); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if modules == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Modules cannot be empty"})
		return
	}

	if len(modules) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Modules cannot be empty"})
		return
	}

	courseId := modules[0].CourseId

	err := controller.moduleService.CreateModules(modules, courseId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Modules created successfully"})
}

func (controller *ModuleController) UpdateModules(context *gin.Context) {
	var modules []*model.Module
	if err := context.ShouldBindJSON(&modules); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if modules == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Modules cannot be empty"})
		return
	}

	err := controller.moduleService.UpdateModules(modules)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Modules updated successfully"})
}

func (controller *ModuleController) SaveBulkOfModules(ctx *gin.Context) {
	var moduleData []dto.ModuleData
	contextApp, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := ctx.ShouldBindJSON(&moduleData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = controller.bulkModuleService.SaveModulesHierarchical(contextApp, moduleData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "error when save")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
