package service

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/repository"
)

type QuizQuestionService interface {
	GetAllQuestionsByQuizId(quizId int64) []model.QuizQuestion

	CreateQuestions(quizId int64, questions []model.QuizQuestion) error

	UpdateQuestions(questions []model.QuizQuestion) error
}

type quizQuestionServiceImpl struct {
	quizQuestionRepository repository.QuizQuestionRepository
}

func (q quizQuestionServiceImpl) GetAllQuestionsByQuizId(quizId int64) []model.QuizQuestion {
	questions, err := q.quizQuestionRepository.GetAllQuizQuestionsByQuizId(quizId)

	if err != nil {
		return nil
	}

	return questions
}

func (q quizQuestionServiceImpl) CreateQuestions(quizId int64, questions []model.QuizQuestion) error {

	return q.quizQuestionRepository.CreateQuizQuestions(questions, quizId)

}

func (q quizQuestionServiceImpl) UpdateQuestions(questions []model.QuizQuestion) error {

	return q.quizQuestionRepository.UpdateQuizQuestions(questions)
}

func NewQuizQuestionService(repo repository.QuizQuestionRepository) QuizQuestionService {
	return &quizQuestionServiceImpl{
		quizQuestionRepository: repo,
	}
}
