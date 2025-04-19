package dto

import "time"

type UserDTO struct {
	UserID    int64     `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	UpdatedBy int64     `json:"updatedBy"`
	CreatedBy int64     `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
