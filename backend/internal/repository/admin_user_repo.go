package repository

import (
	"database/sql"
	"github.com/phamhuy26111995/hto-elearning/internal/database"
	"github.com/phamhuy26111995/hto-elearning/internal/dto"
	"github.com/phamhuy26111995/hto-elearning/internal/model"
)

type AdminUserRepo interface {
	GetAllByRole(role string, paging dto.Paging) ([]model.User, error, int64)
}

type AdminUserRepoImpl struct{}

func NewAdminUserRepository() AdminUserRepo {
	return &AdminUserRepoImpl{}
}

func (repo *AdminUserRepoImpl) GetAllByRole(role string, paging dto.Paging) ([]model.User, error, int64) {
	offset := int((paging.PageNumber - 1) * paging.RowsPerPage)
	limit := int(paging.RowsPerPage)

	query := `SELECT user_id, username, email, role, created_at, updated_at FROM elearning.users WHERE role = $1  ORDER BY created_at DESC 
LIMIT $2 OFFSET $3
`
	countQuery := `SELECT COUNT(*) FROM elearning.users WHERE role = $1`

	rows, err := database.DB.Query(query, role, limit, offset)

	if err != nil {
		return []model.User{}, err, 0
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			panic("Error When Close Row")
		}
	}(rows)

	var list []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return list, err, 0
		}
		list = append(list, user)
	}

	// Thực thi count query
	var totalCount int64
	errCount := database.DB.QueryRow(countQuery, role).Scan(&totalCount)
	if errCount != nil {
		return nil, errCount, 0
	}

	return list, nil, totalCount
}
