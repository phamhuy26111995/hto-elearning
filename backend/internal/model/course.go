package model

import "time"

type Course struct {
	CourseId    int64     `json:"courseId"`
	Title       string    `json:"title"`
	TeacherId   int64     `json:"teacherId" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedBy   int64     `json:"updatedBy"`
	CreatedBy   int64     `json:"createdBy"`
}
