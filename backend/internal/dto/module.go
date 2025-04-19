package dto

type ModuleDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OrderIndex  int    `json:"orderIndex"`
}
