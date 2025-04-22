package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
)

type LessonService interface {
	GetLessonsByModuleId(moduleId int64) ([]model.Lessons, error)
	CreateLessons(lessons []model.Lessons, moduleId int64) error

	UpdateLessons(lessons []model.Lessons) error
}

type lessonServiceImpl struct {
	lessonRepository repository.LessonRepository
}

func (l lessonServiceImpl) GetLessonsByModuleId(moduleId int64) ([]model.Lessons, error) {
	return l.lessonRepository.GetAllLessonsByModuleId(moduleId)
}

func (l lessonServiceImpl) CreateLessons(lessons []model.Lessons, moduleId int64) error {
	return l.lessonRepository.CreateLesson(lessons, moduleId)
}

func (l lessonServiceImpl) UpdateLessons(lessons []model.Lessons) error {
	return l.lessonRepository.UpdateLessons(lessons)
}

func NewLessonService(lessonRepository repository.LessonRepository) LessonService {
	return &lessonServiceImpl{
		lessonRepository: lessonRepository,
	}
}
