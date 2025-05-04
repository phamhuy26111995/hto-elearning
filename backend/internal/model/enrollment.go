package model

import "time"

type Enrollment struct {
	EnrollmentId   int64     `json:"enrollmentId"`
	CourseId       int64     `json:"courseId"`
	UserId         int64     `json:"userId" binding:"required"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
}
