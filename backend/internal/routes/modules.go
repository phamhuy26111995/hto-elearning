package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func RegisterModuleController() *controller.ModuleController {

	moduleRepo := repository.NewModuleRepository()
	moduleService := service.NewModuleService(moduleRepo)
	bulkModuleService := service.NewBulkSaveService(database.DB)
	moduleController := controller.NewModuleController(moduleService, bulkModuleService)

	return &moduleController
}
