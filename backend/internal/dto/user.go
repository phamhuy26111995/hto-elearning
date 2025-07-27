package dto

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"time"
)

type UserDTO struct {
	UserID    int64     `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty"`
	Status    string    `json:"status" default:"ACTIVE"`
	UpdatedBy int64     `json:"updatedBy"`
	CreatedBy int64     `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (userDto UserDTO) MappingToUserEntity(isCreate bool) *model.User {

	user := model.User{
		UserID:    userDto.UserID,
		Username:  userDto.Username,
		Email:     userDto.Email,
		Role:      userDto.Role,
		Status:    userDto.Status,
		Password:  userDto.Password,
		UpdatedBy: userDto.UpdatedBy,
		UpdatedAt: time.Now(),
	}

	if isCreate {
		user.CreatedAt = time.Now()
		user.CreatedBy = userDto.CreatedBy
	}

	return &user
}

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequestDTO struct {
	UserID   int64  `json:"userId"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
