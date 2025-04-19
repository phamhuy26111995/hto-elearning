package model

type Module struct {
	ModuleId    int64  `json:"moduleId"`
	CourseId    int64  `json:"courseId" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OrderIndex  int    `json:"orderIndex"`
}
