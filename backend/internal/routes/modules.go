package routes

import (
	"github.com/phamhuy26111995/hto-elearning/internal/controller"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/service"
)

func registerModuleController() *controller.ModuleController {

	moduleRepo := repository.NewModuleRepository()
	moduleService := service.NewModuleService(moduleRepo)
	moduleController := controller.NewModuleController(moduleService)

	return &moduleController
}
