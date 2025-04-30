package repository

import (
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
)

type QuizOptionRepository interface {
	GetAllQuizOptionsByQuestionId(questionId int64) ([]model.QuizOption, error)
	CreateQuizOptions(quizOptions []model.QuizOption, questionId int64) error
	UpdateQuizOptions(quizOptions []model.QuizOption) error
}

type quizOptionRepositoryImpl struct {
}

func (q quizOptionRepositoryImpl) GetAllQuizOptionsByQuestionId(questionId int64) ([]model.QuizOption, error) {
	query := "SELECT option_id, question_id, option_content, is_correct, order_index FROM elearning.quiz_options WHERE question_id = $1"

	rows, err := database.DB.Query(query, questionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizOptions []model.QuizOption
	for rows.Next() {
		var quizOption model.QuizOption
		err := rows.Scan(&quizOption.OptionId, &quizOption.QuestionId, &quizOption.OptionContent, &quizOption.IsCorrect, &quizOption.OrderIndex)
		if err != nil {
			return nil, err
		}
		quizOptions = append(quizOptions, quizOption)
	}

	return quizOptions, nil
}

func (q quizOptionRepositoryImpl) CreateQuizOptions(quizOptions []model.QuizOption, questionId int64) error {
	if len(quizOptions) == 0 {
		return nil
	}
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, quizOption := range quizOptions {
		query := "INSERT INTO elearning.quiz_options (question_id, option_content, is_correct, order_index) VALUES ($1, $2, $3 , $4)"
		_, err := tx.Exec(query, questionId, quizOption.OptionContent, quizOption.IsCorrect, quizOption.OrderIndex)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (q quizOptionRepositoryImpl) UpdateQuizOptions(quizOptions []model.QuizOption) error {
	if len(quizOptions) == 0 {
		return nil
	}
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, quizOption := range quizOptions {
		query := "UPDATE elearning.quiz_options SET option_content = $1, is_correct = $2, order_index = $4 WHERE option_id = $3"
		_, err := tx.Exec(query, quizOption.OptionContent, quizOption.IsCorrect, quizOption.OptionId, quizOption.OrderIndex)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func NewQuizOptionRepository() QuizOptionRepository {
	return &quizOptionRepositoryImpl{}
}
