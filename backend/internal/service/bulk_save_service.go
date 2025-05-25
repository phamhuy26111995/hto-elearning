package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"strings"
)

type BulkSaveService struct {
	db         *sql.DB
	batchSize  int
	maxWorkers int
}

func NewBulkSaveService(db *sql.DB) BulkSaveService {
	return BulkSaveService{
		db:         db,
		batchSize:  1000, // Số lượng records mỗi batch
		maxWorkers: 5,    // Số lượng workers đồng thời
	}
}

func (s *BulkSaveService) SaveModulesHierarchical(context context.Context, modulesData []dto.ModuleData) error {
	if len(modulesData) == 0 {
		return nil
	}

	tx, err := s.db.BeginTx(context, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction %w", err)
	}

	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				// Nếu rollback thất bại, ghi log lỗi rollback
				fmt.Printf("failed to rollback transaction: %v\n", rollbackErr)
			}
		}
	}()

	// Step 1: Batch insert modules
	moduleIDs, err := s.batchInsertModules(context, tx, modulesData)
	if err != nil {
		return fmt.Errorf("failed to insert modules: %w", err)
	}

	// Step 2: Prepare lessons with module IDs
	var allLessons []model.Lessons
	for i, moduleData := range modulesData {
		for _, lesson := range moduleData.Lessons {
			lesson.ModuleId = moduleIDs[i]
			allLessons = append(allLessons, lesson)
		}
	}

	err = s.batchInsertLessons(context, tx, allLessons)
	if err != nil {
		return fmt.Errorf("failed to insert lessons: %w", err)
	}

	// Step 4: Prepare quizzes with lesson IDs
	var allQuizzes []dto.QuizDto

	for index, moduleData := range modulesData {

		for _, quiz := range moduleData.Quizzes {
			quiz.ModuleId = moduleIDs[index]
			allQuizzes = append(allQuizzes, quiz)
		}
	}

	// Step 5: Batch insert quizzes
	errQuiz := s.batchInsertQuizzes(context, tx, allQuizzes)
	if errQuiz != nil {
		return fmt.Errorf("failed to insert quizzes: %w", err)
	}

	// Step 6: Prepare questions and options
	var allQuestions []dto.QuestionDto
	//var allOptions []model.QuizOption

	for _, quiz := range allQuizzes {

		for _, question := range quiz.Questions {
			question.QuizId = quiz.QuizId
			allQuestions = append(allQuestions, question)
		}

	}

	// Step 7: Batch insert questions
	errQuestion := s.batchInsertQuestions(context, tx, allQuestions)
	if errQuestion != nil {
		return fmt.Errorf("failed to insert questions: %w", err)
	}

	var allOptions []model.QuizOption

	for _, question := range allQuestions {
		for _, option := range question.Options {
			option.QuestionId = question.QuestionID
			allOptions = append(allOptions, option)
		}
	}

	// Step 9: Batch insert options
	if err = s.batchInsertOptions(context, tx, allOptions); err != nil {
		return fmt.Errorf("failed to insert options: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *BulkSaveService) batchInsertModules(ctx context.Context, tx *sql.Tx, modulesData []dto.ModuleData) ([]int64, error) {
	if len(modulesData) == 0 {
		return nil, nil
	}

	// Build bulk insert query với cú pháp PostgreSQL
	query := "INSERT INTO elearning.modules (title, description, course_id, order_index) VALUES "
	values := make([]interface{}, 0, len(modulesData)*2)
	placeholders := make([]string, 0, len(modulesData))
	placeholderIndex := 1

	for _, moduleData := range modulesData {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d , $%d , $%d)", placeholderIndex, placeholderIndex+1, placeholderIndex+2, placeholderIndex+3))
		values = append(values, moduleData.Module.Title, moduleData.Module.Description, moduleData.Module.CourseId, moduleData.Module.OrderIndex)
		placeholderIndex += 4
	}

	query += strings.Join(placeholders, ", ") + " RETURNING module_id"

	rows, err := tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *BulkSaveService) batchInsertLessons(ctx context.Context, tx *sql.Tx, lessons []model.Lessons) error {
	if len(lessons) == 0 {
		return nil
	}

	var allIDs []int64
	for i := 0; i < len(lessons); i += s.batchSize {
		end := i + s.batchSize
		if end > len(lessons) {
			end = len(lessons)
		}

		batch := lessons[i:end]
		ids, err := s.insertLessonBatch(ctx, tx, batch)
		if err != nil {
			return err
		}
		allIDs = append(allIDs, ids...)
	}

	return nil
}

func (s *BulkSaveService) insertLessonBatch(ctx context.Context, tx *sql.Tx, lessons []model.Lessons) ([]int64, error) {
	query := "INSERT INTO elearning.lessons (module_id, title, content, video_url,order_index) VALUES "
	values := make([]interface{}, 0, len(lessons)*4)
	placeholders := make([]string, 0, len(lessons))
	placeholderIndex := 1

	for _, lesson := range lessons {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", placeholderIndex, placeholderIndex+1, placeholderIndex+2, placeholderIndex+3, placeholderIndex+4))
		values = append(values, lesson.ModuleId, lesson.Title, lesson.Content, lesson.VideoUrl, lesson.OrderIndex)
		placeholderIndex += 5
	}

	query += strings.Join(placeholders, ", ") + " RETURNING lesson_id"

	rows, err := tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *BulkSaveService) batchInsertQuizzes(ctx context.Context, tx *sql.Tx, quizzes []dto.QuizDto) error {
	if len(quizzes) == 0 {
		return nil
	}

	query := "INSERT INTO elearning.quizzes (module_id, title, order_index) VALUES "
	values := make([]interface{}, 0, len(quizzes)*3)
	placeholders := make([]string, 0, len(quizzes))
	placeholderIndex := 1

	for _, quiz := range quizzes {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d)", placeholderIndex, placeholderIndex+1, placeholderIndex+2))
		values = append(values, quiz.ModuleId, quiz.Title, quiz.OrderIndex)
		placeholderIndex += 3
	}

	query += strings.Join(placeholders, ", ") + " RETURNING quiz_id"

	rows, err := tx.QueryContext(ctx, query, values...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for index, quiz := range quizzes {
		quiz.QuizId = ids[index]
		quizzes[index] = quiz
	}

	return nil
}

func (s *BulkSaveService) batchInsertQuestions(ctx context.Context, tx *sql.Tx, questions []dto.QuestionDto) error {
	if len(questions) == 0 {
		return nil
	}

	query := "INSERT INTO elearning.quiz_questions (quiz_id, question_content, question_type, order_index) VALUES "
	values := make([]interface{}, 0, len(questions)*4)
	placeholders := make([]string, 0, len(questions))
	placeholderIndex := 1

	for _, question := range questions {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", placeholderIndex, placeholderIndex+1, placeholderIndex+2, placeholderIndex+3))
		values = append(values, question.QuizId, question.QuestionContent, question.QuestionType, question.OrderIndex)
		placeholderIndex += 4
	}

	query += strings.Join(placeholders, ", ") + " RETURNING question_id"

	rows, err := tx.QueryContext(ctx, query, values...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for index, question := range questions {
		question.QuestionID = ids[index]
		questions[index] = question
	}

	return nil
}

func (s *BulkSaveService) batchInsertOptions(ctx context.Context, tx *sql.Tx, options []model.QuizOption) error {
	if len(options) == 0 {
		return nil
	}

	query := "INSERT INTO elearning.quiz_options (question_id, option_content, is_correct, order_index) VALUES "
	values := make([]interface{}, 0, len(options)*4)
	placeholders := make([]string, 0, len(options))
	placeholderIndex := 1

	for _, option := range options {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", placeholderIndex, placeholderIndex+1, placeholderIndex+2, placeholderIndex+3))
		values = append(values, option.QuestionId, option.OptionContent, option.IsCorrect, option.OrderIndex)
		placeholderIndex += 4
	}

	query += strings.Join(placeholders, ", ")

	_, err := tx.ExecContext(ctx, query, values...)
	return err
}
