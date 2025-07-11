package service

import (
	"errors"
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
	"github.com/phamhuy26111995/hto-elearning/internal/utils"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserById(userId int64) (*model.User, error)

	GetAllUsersByTeacherId(teacherId int64) ([]dto.UserDTO, error)

	CreateUser(user *model.User) error
	CreateStudent(user *model.User, teacherId int64) error

	UpdateUser(user *model.User) error

	DeleteUserPermanently(userId int64, teacherId int64) error

	Login(user *dto.UserLoginDTO) (jwt string, userInfo *model.User)

	EnrollCourseForStudent(userId int64, courseId int64) error

	UnEnrollCourseForStudent(userId int64, courseId int64) error

	ChangeStatus(studentId int64, status string) error
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func (service *userServiceImpl) DeleteUserPermanently(userId int64, teacherId int64) error {
	isValid, err := service.repo.CheckValidToDelete(userId, teacherId)

	if !isValid {
		if err != nil {
			return err
		} else {
			return errors.New("Invalid User To Delete")
		}
	}

	err = service.repo.DeleteUser(userId)

	return err

}

func (service *userServiceImpl) ChangeStatus(studentId int64, status string) error {
	return service.repo.ChangeStatus(studentId, status)
}

func (service *userServiceImpl) EnrollCourseForStudent(userId int64, courseId int64) error {
	err := service.repo.EnrollCourseForStudent(userId, courseId)
	if err != nil {
		return err
	}

	return nil
}

func (service *userServiceImpl) UnEnrollCourseForStudent(userId int64, courseId int64) error {
	err := service.repo.UnEnrollCourseForStudent(userId, courseId)
	if err != nil {
		return err
	}

	return nil
}

func (service *userServiceImpl) CreateStudent(user *model.User, teacherId int64) error {
	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword
	return service.repo.CreateStudent(user, teacherId)
}

func (service *userServiceImpl) GetAllUsersByTeacherId(teacherId int64) ([]dto.UserDTO, error) {
	models, err := service.repo.GetAllByTeacherId(teacherId)

	if err != nil {
		return nil, err
	}

	dtos := make([]dto.UserDTO, len(models))

	for i, m := range models {
		dtos[i] = dto.UserDTO{
			UserID:    m.UserID,
			Username:  m.Username,
			Email:     m.Email,
			Role:      m.Role,
			CreatedAt: m.CreatedAt,
			CreatedBy: m.CreatedBy,
			UpdatedAt: m.UpdatedAt,
			UpdatedBy: m.UpdatedBy,
		}
	}

	// 4. Return the DTO slice
	return dtos, nil
}

func (service *userServiceImpl) Login(user *dto.UserLoginDTO) (jwt string, userInfo *model.User) {
	queryUser, err := service.repo.GetUserByUsernameToVal(user.Username)

	if err != nil {
		return "", nil
	}

	if utils.CheckPasswordHash(user.Password, queryUser.Password) {
		token, _ := utils.GenerateToken(queryUser.Username, queryUser.UserID, queryUser.Role)
		return token, queryUser
	}

	return "", nil
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
