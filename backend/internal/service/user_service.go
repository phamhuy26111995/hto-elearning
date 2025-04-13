package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/utils"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(userId int64) (*model.User, error)

	CreateUser(user *model.User) error

	UpdateUser(user *model.User) error
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func (service *userServiceImpl) GetUserById(userId int64) (*model.User, error) {
	return service.repo.GetUserById(userId)
}

func (service *userServiceImpl) CreateUser(user *model.User) error {
	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword
	return service.repo.CreateUser(user)
}

func (service *userServiceImpl) UpdateUser(user *model.User) error {
	return service.repo.UpdateUser(user)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (service *userServiceImpl) GetAllUsers() ([]model.User, error) {
	return service.repo.GetAll()
}
