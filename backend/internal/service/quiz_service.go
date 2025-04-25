package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
)

type QuizService interface {
	CreateQuizzes(quizzes []model.Quiz) error
	GetAllQuizzesByModuleId(moduleId int64) ([]model.Quiz, error)
	UpdateQuizzes(quizzes []model.Quiz) error
}

type quizServiceImpl struct {
	quizRepository repository.QuizRepository
}

func (q quizServiceImpl) CreateQuizzes(quizzes []model.Quiz) error {

	if len(quizzes) == 0 {
		return nil
	}

	return q.quizRepository.CreateQuizzes(quizzes, quizzes[0].ModuleId)
}

func (q quizServiceImpl) GetAllQuizzesByModuleId(moduleId int64) ([]model.Quiz, error) {
	return q.quizRepository.GetAllQuizzesByModuleId(moduleId)
}

func (q quizServiceImpl) UpdateQuizzes(quizzes []model.Quiz) error {
	if len(quizzes) == 0 {
		return nil
	}
	return q.quizRepository.UpdateQuizzes(quizzes)
}

func NewQuizService(quizRepository repository.QuizRepository) QuizService {
	return &quizServiceImpl{
		quizRepository: quizRepository,
	}
}
