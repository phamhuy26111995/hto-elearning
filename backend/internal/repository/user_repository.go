package repository

import (
	"errors"
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"strings"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	CreateUser(user *model.User) error
	GetUserById(userId int64) (*model.User, error)
	UpdateUser(user *model.User) error
}

type userRepositoryImpl struct {
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
		setParts = append(setParts, fmt.Sprintf("password = $%d", placeholderIndex))
		args = append(args, user.Password)
		placeholderIndex++
	}
	// Cột updated_by kiểu số, nếu khác 0 thì cập nhật.
	if user.UpdatedBy != 0 {
		setParts = append(setParts, fmt.Sprintf("updated_by = $%d", placeholderIndex))
		args = append(args, user.UpdatedBy)
		placeholderIndex++
	}
	// Cột created_by kiểu số, nếu khác 0 thì cập nhật.
	if user.CreatedBy != 0 {
		setParts = append(setParts, fmt.Sprintf("created_by = $%d", placeholderIndex))
		args = append(args, user.CreatedBy)
		placeholderIndex++
	}

	// Nếu không có trường nào có giá trị cập nhật thì báo lỗi.
	if len(setParts) == 0 {
		return errors.New("không có trường nào được cập nhật")
	}

	// Xây dựng câu query với các placeholder theo kiểu PostgreSQL.
	query := fmt.Sprintf("UPDATE users SET %s WHERE user_id = $%d", strings.Join(setParts, ", "), placeholderIndex)
	// Thêm điều kiện WHERE, ở đây cập nhật dựa trên user_id.
	args = append(args, user.UserID)

	// Thực thi câu query.
	_, err := database.DB.Exec(query, args...)
	return err
}

func (u *userRepositoryImpl) GetUserById(userId int64) (*model.User, error) {
	query := "SELECT user_id, username, email, role, created_at, updated_at FROM users WHERE user_id = $1"
	row := database.DB.QueryRow(query, userId)
	var user model.User
	err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepositoryImpl) CreateUser(user *model.User) error {
	query := `INSERT INTO users (username,email, password, role, created_by,updated_by) 
			VALUES ($1, $2, $3, $4, $5, $6)`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.Role, user.CreatedBy, user.UpdatedBy)
	return err
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) GetAll() ([]model.User, error) {
	query := "SELECT user_id, username, email, role, created_at, updated_at FROM users"
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
