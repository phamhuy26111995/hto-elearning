package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
)

type QuizOptionService interface {
	GetAllQuizOptionsByQuestionId(questionId int64) ([]model.QuizOption, error)
	CreateQuizOptions(quizOptions []model.QuizOption, questionId int64) error
	UpdateQuizOptions(quizOptions []model.QuizOption) error
}

type quizOptionServiceImpl struct {
	repo repository.QuizOptionRepository
}

func (q quizOptionServiceImpl) GetAllQuizOptionsByQuestionId(questionId int64) ([]model.QuizOption, error) {
	return q.repo.GetAllQuizOptionsByQuestionId(questionId)
}

func (q quizOptionServiceImpl) CreateQuizOptions(quizOptions []model.QuizOption, questionId int64) error {
	return q.repo.CreateQuizOptions(quizOptions, questionId)
}

func (q quizOptionServiceImpl) UpdateQuizOptions(quizOptions []model.QuizOption) error {
	return q.repo.UpdateQuizOptions(quizOptions)
}

func NewQuizOptionService(repo repository.QuizOptionRepository) QuizOptionService {
	return &quizOptionServiceImpl{
		repo: repo,
	}
}
