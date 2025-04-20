package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
)

type ModuleService interface {
	GetAllModulesByCourse(courseId int64) ([]*dto.ModuleDTO, error)

	CreateModules(modules []*model.Module, courseId int64) error

	UpdateModules(modules []*model.Module) error
}

type moduleServiceImpl struct {
	repo repository.ModuleRepository
}

func (service *moduleServiceImpl) UpdateModules(modules []*model.Module) error {
	return service.repo.UpdateModules(modules)
}

func (service *moduleServiceImpl) CreateModules(modules []*model.Module, courseId int64) error {
	return service.repo.CreateModules(modules, courseId)
}

func (service *moduleServiceImpl) GetAllModulesByCourse(courseId int64) ([]*dto.ModuleDTO, error) {
	modules, err := service.repo.GetAllModulesByCourse(courseId)

	if err != nil {
		return nil, err
	}

	moduleDtoList := make([]*dto.ModuleDTO, len(modules))
	for i, module := range modules {
		moduleDtoList[i] = &dto.ModuleDTO{
			ID:          module.ModuleId,
			Title:       module.Title,
			Description: module.Description,
			OrderIndex:  module.OrderIndex,
		}
	}
	return moduleDtoList, nil

}

func NewModuleService(repo repository.ModuleRepository) ModuleService {
	return &moduleServiceImpl{repo: repo}
}
