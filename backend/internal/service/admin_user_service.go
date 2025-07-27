package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/utils"
)

type AdminUserService interface {
	GetAllByRole(role string, paging dto.Paging) ([]dto.UserDTO, error)

	CreateUser(user *model.User) error
}

type AdminUserServiceImpl struct {
	adminRepo repository.AdminUserRepo

	repo repository.UserRepository
}

func NewAdminUserService(adminRepo repository.AdminUserRepo, repo repository.UserRepository) AdminUserService {
	return &AdminUserServiceImpl{adminRepo: adminRepo, repo: repo}
}

func (service *AdminUserServiceImpl) GetAllByRole(role string, paging dto.Paging) ([]dto.UserDTO, error) {
	userEntities, err := service.adminRepo.GetAllByRole(role, paging)

	if err != nil {
		return []dto.UserDTO{}, err
	}

	dtoList := make([]dto.UserDTO, len(userEntities))

	for index, entity := range userEntities {
		dtoList[index] = dto.UserDTO{
			UserID:    entity.UserID,
			Username:  entity.Username,
			Email:     entity.Email,
			Role:      entity.Role,
			CreatedAt: entity.CreatedAt,
			CreatedBy: entity.CreatedBy,
			UpdatedAt: entity.UpdatedAt,
			UpdatedBy: entity.UpdatedBy,
		}
	}

	return dtoList, nil

}

func (service *AdminUserServiceImpl) CreateUser(user *model.User) error {
	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword
	return service.repo.CreateUser(user)
}
