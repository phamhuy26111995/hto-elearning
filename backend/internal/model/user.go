package model

import "time"

type User struct {
	UserID    int64     `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	UpdatedBy int64     `json:"updatedBy"`
	CreatedBy int64     `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ParentID  int64     `json:"parentID"`
}
