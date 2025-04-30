package repository

import (
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
)

type QuizRepository interface {
	GetAllQuizzesByModuleId(moduleId int64) ([]model.Quiz, error)

	CreateQuizzes(quizzes []model.Quiz, moduleId int64) error

	UpdateQuizzes(quizzes []model.Quiz) error
}

type quizRepositoryImpl struct {
}

func NewQuizRepository() QuizRepository {
	return &quizRepositoryImpl{}
}

func (q *quizRepositoryImpl) GetAllQuizzesByModuleId(moduleId int64) ([]model.Quiz, error) {
	query := `SELECT quiz_id, title, order_index FROM elearning.quizzes WHERE module_id = $1`

	rows, err := database.DB.Query(query, moduleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []model.Quiz
	for rows.Next() {
		var quiz model.Quiz
		err := rows.Scan(&quiz.QuizId, &quiz.Title, &quiz.OrderIndex)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (q *quizRepositoryImpl) CreateQuizzes(quizzes []model.Quiz, moduleId int64) error {
	if len(quizzes) == 0 {
		return nil
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	for _, quiz := range quizzes {
		query := `INSERT INTO elearning.quizzes (module_id, title, order_index) VALUES ($1, $2,$3)`
		_, err := tx.Exec(query, moduleId, quiz.Title, quiz.OrderIndex)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *quizRepositoryImpl) UpdateQuizzes(quizzes []model.Quiz) error {
	if len(quizzes) == 0 {
		return nil
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	for _, quiz := range quizzes {
		query := `UPDATE elearning.quizzes SET title = $1, order_index = $3 WHERE quiz_id = $2`
		_, err := tx.Exec(query, quiz.Title, quiz.QuizId, quiz.OrderIndex)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
