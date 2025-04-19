package repository

import (
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"strings"
)

type ModuleRepository interface {
	GetAllModulesByCourse(courseId int64) ([]*model.Module, error)

	CreateModules(modules []*model.Module, courseId int64) error
}

type moduleRepositoryImpl struct{}

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
		"INSERT INTO modules (title, description, order_index, course_id) VALUES %s",
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
	SELECT module_id, title, description, order_index FROM modules WHERE course_id = $1
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
