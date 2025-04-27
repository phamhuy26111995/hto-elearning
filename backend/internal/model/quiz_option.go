package model

type QuizOption struct {
	OptionId      int64  `json:"optionId"`
	QuestionId    int64  `json:"questionId"`
	OptionContent string `json:"optionContent"`
	IsCorrect     bool   `json:"isCorrect"`
	OrderIndex    int    `json:"orderIndex"`
}
