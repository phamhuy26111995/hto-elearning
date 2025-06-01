package model

import "time"

type QuizQuestion struct {
	QuestionID      int64     `json:"questionId"`
	QuizId          int64     `json:"quizId,omitempty"`
	QuestionContent string    `json:"questionContent"`
	QuestionType    string    `json:"questionType"`
	OrderIndex      int       `json:"orderIndex"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
