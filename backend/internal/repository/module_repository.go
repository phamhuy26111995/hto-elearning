package repository

import (
	"errors"
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"strings"
)

type ModuleRepository interface {
	GetAllModulesByCourse(courseId int64) ([]*model.Module, error)

	CreateModules(modules []*model.Module, courseId int64) error

	UpdateModules(modules []*model.Module) error
}

type moduleRepositoryImpl struct{}

func (m *moduleRepositoryImpl) UpdateModules(modules []*model.Module) error {
	// Kiểm tra danh sách modules
	if len(modules) == 0 {
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
	for _, module := range modules {
		// Kiểm tra moduleId
		if module.ModuleId <= 0 {
			tx.Rollback()
			return errors.New("moduleId không hợp lệ")
		}

		// Tạo câu truy vấn động chỉ cho các trường có giá trị
		query := "UPDATE elearning.modules SET "
		params := []interface{}{}
		updateFields := []string{}
		paramIndex := 1 // PostgreSQL sử dụng $1, $2, ... thay vì ?

		if module.Title != "" {
			updateFields = append(updateFields, fmt.Sprintf("title = $%d", paramIndex))
			params = append(params, module.Title)
			paramIndex++
		}

		if module.Description != "" {
			updateFields = append(updateFields, fmt.Sprintf("description = $%d", paramIndex))
			params = append(params, module.Description)
			paramIndex++
		}

		if module.OrderIndex != 0 {
			updateFields = append(updateFields, fmt.Sprintf("order_index = $%d", paramIndex))
			params = append(params, module.OrderIndex)
			paramIndex++
		}

		// Nếu không có trường nào cần cập nhật, bỏ qua module này
		if len(updateFields) == 0 {
			continue
		}

		// Hoàn thành câu truy vấn
		query += strings.Join(updateFields, ", ")
		query += fmt.Sprintf(" WHERE module_id = $%d", paramIndex)
		params = append(params, module.ModuleId)

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

func (m *moduleRepositoryImpl) CreateModules(modules []*model.Module, courseId int64) error {
	if len(modules) == 0 {
		return nil
	}
	var placeholders []string
	var args []interface{}

	for i, module := range modules {
		idx := i * 4
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", idx+1, idx+2, idx+3, idx+4))
		args = append(args, module.Title, module.Description, module.OrderIndex, courseId)
	}

	query := fmt.Sprintf(
		"INSERT INTO elearning.modules (title, description, order_index, course_id) VALUES %s",
		strings.Join(placeholders, ","),
	)

	stmt, _ := database.DB.Prepare(query)
	defer stmt.Close()

	_, err := stmt.Exec(args...)
	return err
}

func NewModuleRepository() ModuleRepository {
	return &moduleRepositoryImpl{}
}

func (m *moduleRepositoryImpl) GetAllModulesByCourse(courseId int64) ([]*model.Module, error) {
	query := `
	SELECT module_id, title, description, order_index FROM elearning.modules WHERE course_id = $1
`
	rows, err := database.DB.Query(query, courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []*model.Module
	for rows.Next() {
		var module model.Module
		err := rows.Scan(&module.CourseId, &module.Title, &module.Description)
		if err != nil {
			return nil, err
		}
		modules = append(modules, &module)
	}

	return modules, nil
}
