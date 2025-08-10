package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"github.com/phamhuy26111995/hto-elearning/internal/utils"
	"strings"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	CreateUser(user *model.User) error

	CreateStudent(student *model.User, teacherId int64) error

	GetUserById(userId int64) (*model.User, error)
	UpdateUser(user *model.User) error

	GetUserByUsernameToVal(username string) (*model.User, error)

	GetAllByTeacherId(teacherId int64) ([]model.User, error)

	EnrollCourseForStudent(studentId int64, courseId int64) error
	UnEnrollCourseForStudent(studentId int64, courseId int64) error

	ChangeStatus(studentId int64, status string) error

	DeleteUser(userId int64) error

	CheckValidToDelete(userId int64, teacherId int64) (bool, error)
}

type userRepositoryImpl struct {
}

func (u *userRepositoryImpl) DeleteUser(userId int64) error {
	query := `DELETE FROM elearning.users WHERE user_id = $1`
	_, err := database.DB.Exec(query, userId)

	return err
}

func (u *userRepositoryImpl) CheckValidToDelete(userId int64, teacherId int64) (bool, error) {
	query := `SELECT 1 FROM elearning.users WHERE user_id = $1 AND teacher_id = $2`
	var exists int
	err := database.DB.QueryRow(query, userId, teacherId).Scan(&exists)
	if err == sql.ErrNoRows {
		// No matching row → userId is not valid under that teacherId
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userRepositoryImpl) UnEnrollCourseForStudent(studentId int64, courseId int64) error {
	query := `DELETE FROM elearning.enrollments WHERE user_id = $1 AND course_id = $2`
	_, err := database.DB.Exec(query, studentId, courseId)
	return err
}

func (u *userRepositoryImpl) ChangeStatus(studentId int64, status string) error {
	query := `UPDATE elearning.users SET status = $1 WHERE user_id = $2 AND role = 'STUDENT'`
	_, err := database.DB.Exec(query, status, studentId)
	return err
}

func (u *userRepositoryImpl) EnrollCourseForStudent(studentId int64, courseId int64) error {
	query := `INSERT INTO elearning.enrollments (user_id, course_id) VALUES ($1, $2)`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(studentId, courseId)
	return err
}

func (u *userRepositoryImpl) CreateStudent(student *model.User, teacherId int64) error {
	query := `INSERT INTO elearning.users (username,email, password, created_by,updated_by,teacher_id, created_at , updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, NOW() , NOW())`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(student.Username, student.Email, student.Password, teacherId, teacherId, teacherId)
	return err
}

func (u *userRepositoryImpl) GetAllByTeacherId(teacherId int64) ([]model.User, error) {
	query := `SELECT user_id, username, email, role, created_at, updated_at, created_by, updated_by FROM elearning.users WHERE teacher_id = $1`
	rows, err := database.DB.Query(query, teacherId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepositoryImpl) GetUserByUsernameToVal(username string) (*model.User, error) {
	var user model.User
	query := `SELECT user_id,username,email,password, role  FROM elearning.users WHERE username = $1`
	err := database.DB.QueryRow(query, username).Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) UpdateUser(user *model.User) error {
	var setParts []string
	var args []interface{}
	placeholderIndex := 1

	// Kiểm tra các trường cần cập nhật, nếu có giá trị thì thêm vào câu query và danh sách tham số.
	if user.Username != "" {
		setParts = append(setParts, fmt.Sprintf("username = $%d", placeholderIndex))
		args = append(args, user.Username)
		placeholderIndex++
	}
	if user.Email != "" {
		setParts = append(setParts, fmt.Sprintf("email = $%d", placeholderIndex))
		args = append(args, user.Email)
		placeholderIndex++
	}
	if user.Password != "" {
		hashPassword, _ := utils.HashPassword(user.Password)
		user.Password = hashPassword
		setParts = append(setParts, fmt.Sprintf("password = $%d", placeholderIndex))
		args = append(args, user.Password)
		placeholderIndex++
	}

	if user.Status != "" {
		setParts = append(setParts, fmt.Sprintf("status = $%d", placeholderIndex))
		args = append(args, user.Status)
		placeholderIndex++
	}

	setParts = append(setParts, fmt.Sprintf("updated_by = $%d", placeholderIndex))
	args = append(args, user.UpdatedBy)
	placeholderIndex++

	setParts = append(setParts, "updated_at = NOW() ")

	// Nếu không có trường nào có giá trị cập nhật thì báo lỗi.
	if len(setParts) == 0 {
		return errors.New("không có trường nào được cập nhật")
	}

	// Xây dựng câu query với các placeholder theo kiểu PostgreSQL.
	query := fmt.Sprintf("UPDATE elearning.users SET %s WHERE user_id = $%d", strings.Join(setParts, ", "), placeholderIndex)
	// Thêm điều kiện WHERE, ở đây cập nhật dựa trên user_id.
	args = append(args, user.UserID)

	// Thực thi câu query.
	_, err := database.DB.Exec(query, args...)
	return err
}

func (u *userRepositoryImpl) GetUserById(userId int64) (*model.User, error) {
	query := "SELECT user_id, username, email, role, created_at, updated_at FROM elearning.users WHERE user_id = $1"
	row := database.DB.QueryRow(query, userId)
	var user model.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepositoryImpl) CreateUser(user *model.User) error {
	tx, txerr := database.DB.Begin()
	if txerr != nil {
		return txerr
	}
	defer tx.Rollback()

	query := `INSERT INTO elearning.users (username,email, password, role, created_by,updated_by, status, parent_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7,$8) RETURNING user_id`
	err := database.DB.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedBy,
		user.UpdatedBy,
		user.Status,
		user.ParentID,
	).Scan(&user.UserID)

	if err != nil {
		return err
	}
	tx.Commit()

	//_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Role, user.CreatedBy, user.UpdatedBy, user.Status)
	return err
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) GetAll() ([]model.User, error) {
	query := "SELECT user_id, username, email, role, created_at, updated_at FROM elearning.users"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
