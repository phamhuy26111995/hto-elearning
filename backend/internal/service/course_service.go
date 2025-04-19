package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
)

type CourseService interface {
	GetAllCourses(userId int64) ([]model.Course, error)
	GetCourse(id int64) (*model.Course, error)
	CreateCourse(course *model.Course, userId int64) error
	UpdateCourse(course *model.Course) error
	DeleteCourse(id int64) error
}

type courseServiceImpl struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseServiceImpl{repo: repo}
}

func (service *courseServiceImpl) GetAllCourses(userId int64) ([]model.Course, error) {

	return service.repo.GetAllCoursesByUserId(userId)
}

func (service *courseServiceImpl) GetCourse(id int64) (*model.Course, error) {
	return service.repo.GetCourseById(id)
}

func (service *courseServiceImpl) CreateCourse(course *model.Course, userId int64) error {
	return service.repo.CreateCourse(course, userId)
}

func (service *courseServiceImpl) UpdateCourse(course *model.Course) error {
	return service.repo.UpdateCourse(course)
}

func (service *courseServiceImpl) DeleteCourse(id int64) error {
	return service.repo.DeleteCourse(id)
}
