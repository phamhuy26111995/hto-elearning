package repository

import (
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"strings"
)

type CourseRepository interface {
	GetAllCoursesByUserId(userId int64) ([]model.Course, error)
	CreateCourse(course *model.Course, userId int64) error

	GetCourseById(courseId int64) (*model.Course, error)

	UpdateCourse(course *model.Course) error

	DeleteCourse(courseId int64) error
}

type courseRepositoryImpl struct{}

func NewCourseRepository() CourseRepository {
	return &courseRepositoryImpl{}
}

func (c *courseRepositoryImpl) GetAllCoursesByUserId(userId int64) ([]model.Course, error) {
	query := `
	SELECT course_id, title, description FROM elearning.courses WHERE teacher_id = $1
`
	rows, err := database.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Course
	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *courseRepositoryImpl) CreateCourse(course *model.Course, userId int64) error {
	query := `INSERT INTO elearning.courses (title, description, teacher_id, created_by, updated_by) 
			VALUES ($1, $2, $3, $4, $5)`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(course.Title, course.Description, userId, course.CreatedBy, course.UpdatedBy)
	return err
}

func (c *courseRepositoryImpl) GetCourseById(courseId int64) (*model.Course, error) {
	query := "SELECT course_id, title, description, teacher_id FROM elearning.courses WHERE course_id = $1"
	row := database.DB.QueryRow(query, courseId)
	var course model.Course
	err := row.Scan(&course.CourseId, &course.Title, &course.Description, &course.TeacherId)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *courseRepositoryImpl) UpdateCourse(course *model.Course) error {
	var setParts []string
	var args []interface{}
	placeholderIndex := 1

	if course.Title != "" {
		setParts = append(setParts, fmt.Sprintf("title = $%d", placeholderIndex))
		args = append(args, course.Title)
		placeholderIndex++
	}

	if course.Description != "" {
		setParts = append(setParts, fmt.Sprintf("description = $%d", placeholderIndex))
		args = append(args, course.Description)
		placeholderIndex++
	}

	if course.UpdatedBy != 0 {
		setParts = append(setParts, fmt.Sprintf("updated_by = $%d", placeholderIndex))
		args = append(args, course.UpdatedBy)
	}

	query := fmt.Sprintf("UPDATE elearning.courses SET %s WHERE course_id = $%d", strings.Join(setParts, ", "), placeholderIndex)
	args = append(args, course.CourseId)

	_, err := database.DB.Exec(query, args...)
	return err
}

func (c *courseRepositoryImpl) DeleteCourse(courseId int64) error {
	query := "DELETE FROM elearning.courses WHERE course_id = $1"
	_, err := database.DB.Exec(query, courseId)
	return err
}
