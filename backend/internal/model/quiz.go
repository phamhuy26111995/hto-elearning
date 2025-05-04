package model

import "time"

type Quiz struct {
	QuizId     int64     `json:"quizId"`
	ModuleId   int64     `json:"moduleId,omitempty"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"createdAt"`
	OrderIndex int       `json:"orderIndex"`
}
