package repository

import (
	"errors"
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"strings"
)

type LessonRepository interface {
	GetAllLessonsByModuleId(moduleId int64) ([]model.Lessons, error)

	CreateLesson(lessons []model.Lessons, moduleId int64) error

	UpdateLessons(lessons []model.Lessons) error
}

type lessonRepositoryImpl struct{}

func (l lessonRepositoryImpl) GetAllLessonsByModuleId(moduleId int64) ([]model.Lessons, error) {
	query := `SELECT lesson_id, title, content, video_url, order_index FROM lessons WHERE module_id = $1`

	rows, err := database.DB.Query(query, moduleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(&lesson.LessonId, &lesson.Title, &lesson.Content, &lesson.VideoUrl, &lesson.OrderIndex)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

func (l lessonRepositoryImpl) CreateLesson(lessons []model.Lessons, moduleId int64) error {
	if len(lessons) == 0 {
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

	for _, lesson := range lessons {
		query := `INSERT INTO lessons (module_id, title, content, video_url, order_index) VALUES ($1, $2, $3, $4, $5)`
		_, err := tx.Exec(query, moduleId, lesson.Title, lesson.Content, lesson.VideoUrl, lesson.OrderIndex)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (l lessonRepositoryImpl) UpdateLessons(lessons []model.Lessons) error {
	if len(lessons) == 0 {
		return nil // Không có module nào để cập nhật
	}

	// Bắt đầu giao dịch
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

	// Lặp qua từng module và cập nhật
	for _, lesson := range lessons {
		// Kiểm tra moduleId
		if lesson.LessonId <= 0 {
			tx.Rollback()
			return errors.New("moduleId không hợp lệ")
		}

		// Tạo câu truy vấn động chỉ cho các trường có giá trị
		query := "UPDATE lessons SET "
		params := []interface{}{}
		updateFields := []string{}
		paramIndex := 1 // PostgreSQL sử dụng $1, $2, ... thay vì ?

		if lesson.Title != "" {
			updateFields = append(updateFields, fmt.Sprintf("title = $%d", paramIndex))
			params = append(params, lesson.Title)
			paramIndex++
		}

		if lesson.Content != "" {
			updateFields = append(updateFields, fmt.Sprintf("content = $%d", paramIndex))
			params = append(params, lesson.Content)
			paramIndex++
		}

		if lesson.VideoUrl != "" {
			updateFields = append(updateFields, fmt.Sprintf("video_url = $%d", paramIndex))
			params = append(params, lesson.VideoUrl)
			paramIndex++
		}

		if lesson.OrderIndex != 0 {
			updateFields = append(updateFields, fmt.Sprintf("order_index = $%d", paramIndex))
			params = append(params, lesson.OrderIndex)
			paramIndex++
		}

		// Nếu không có trường nào cần cập nhật, bỏ qua module này
		if len(updateFields) == 0 {
			continue
		}

		// Hoàn thành câu truy vấn
		query += strings.Join(updateFields, ", ")
		query += fmt.Sprintf(" WHERE lesson_id = $%d", paramIndex)
		params = append(params, lesson.LessonId)

		// Thực hiện cập nhật
		_, err := tx.Exec(query, params...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit giao dịch
	return tx.Commit()
}

func NewLessonRepository() LessonRepository {
	return &lessonRepositoryImpl{}
}
