package model

type QuizQuestion struct {
	QuestionID      int64  `json:"questionId"`
	QuizId          int64  `json:"quizId"`
	QuestionContent string `json:"questionContent"`
	QuestionType    string `json:"questionType"`
}
