package dto

import (
	"github.com/phamhuy26111995/hto-elearning/internal/model"
	"time"
)

type QuizDto struct {
	QuizId     int64         `json:"quizId"`
	ModuleId   int64         `json:"moduleId,omitempty"`
	Title      string        `json:"title"`
	CreatedAt  time.Time     `json:"createdAt"`
	OrderIndex int           `json:"orderIndex"`
	Questions  []QuestionDto `json:"questions"`
}

type QuestionDto struct {
	QuestionID      int64              `json:"questionId"`
	QuizId          int64              `json:"quizId,omitempty"`
	QuestionContent string             `json:"questionContent"`
	QuestionType    string             `json:"questionType"`
	OrderIndex      int                `json:"orderIndex"`
	Options         []model.QuizOption `json:"options"`
}

type ModuleData struct {
	Module  model.Module    `json:"module"`
	Lessons []model.Lessons `json:"lessons"`
	Quizzes []QuizDto       `json:"quizzes"`
}
