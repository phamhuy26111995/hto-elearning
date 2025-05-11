package dto

import "time"

type EnrollmentDTO struct {
	EnrollmentId   int64     `json:"enrollmentId"`
	CourseId       int64     `json:"courseId" binding:"required"`
	UserId         int64     `json:"userId" binding:"required"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
}
