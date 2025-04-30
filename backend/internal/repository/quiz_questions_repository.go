package repository

import (
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
)

type QuizQuestionRepository interface {
	CreateQuizQuestions(quizQuestions []model.QuizQuestion, quizId int64) error
	GetAllQuizQuestionsByQuizId(quizId int64) ([]model.QuizQuestion, error)
	UpdateQuizQuestions(quizQuestions []model.QuizQuestion) error
}

type quizQuestionRepositoryImpl struct {
}

func (q quizQuestionRepositoryImpl) CreateQuizQuestions(quizQuestions []model.QuizQuestion, quizId int64) error {
	if len(quizQuestions) == 0 {
		return nil
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, quizQuestion := range quizQuestions {
		query := `INSERT INTO elearning.quiz_questions (quiz_id, question_content, question_type,order_index) VALUES ($1, $2, $3, $4)`
		_, err := tx.Exec(query, quizId, quizQuestion.QuestionContent, quizQuestion.QuestionType, quizQuestion.OrderIndex)
		if err != nil {
			return err
		}
	}
	return tx.Commit()

}

func (q quizQuestionRepositoryImpl) GetAllQuizQuestionsByQuizId(quizId int64) ([]model.QuizQuestion, error) {
	query := `SELECT question_id ,quiz_id, question_content, question_type, order_index FROM elearning.quiz_questions WHERE quiz_id = $1`

	rows, err := database.DB.Query(query, quizId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizQuestions []model.QuizQuestion
	for rows.Next() {
		var quizQuestion model.QuizQuestion
		err := rows.Scan(&quizQuestion.QuestionID, &quizQuestion.QuizId, &quizQuestion.QuestionContent, &quizQuestion.QuestionType, &quizQuestion.OrderIndex)
		if err != nil {
			return nil, err
		}
		quizQuestions = append(quizQuestions, quizQuestion)
	}
	return quizQuestions, nil
}

func (q quizQuestionRepositoryImpl) UpdateQuizQuestions(quizQuestions []model.QuizQuestion) error {
	if len(quizQuestions) == 0 {
		return nil
	}
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, quizQuestion := range quizQuestions {
		query := `UPDATE elearning.quiz_questions SET question_content = $1, question_type = $2, order_index = $4 WHERE question_id = $3`
		_, err := tx.Exec(query, quizQuestion.QuestionContent, quizQuestion.QuestionType, quizQuestion.QuestionID, quizQuestion.OrderIndex)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func NewQuizQuestionRepository() QuizQuestionRepository {
	return &quizQuestionRepositoryImpl{}
}
