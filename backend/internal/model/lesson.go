package model

import "time"

type Lessons struct {
	LessonId   int64     `json:"lessonId"`
	ModuleId   int64     `json:"moduleId,omitempty"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	VideoUrl   string    `json:"videoUrl"`
	OrderIndex int       `json:"orderIndex"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
