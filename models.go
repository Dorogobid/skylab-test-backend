package main

type Category struct {
	ID      uint    `json:"id" gorm:"primaryKey" example:"1"`
	Name	string	`json:"name" example:"Category name"`
}

type CategoryCreate struct {
	Name	string	` example:"Category name"`
}

type Question struct {
	ID      		uint    `json:"id" gorm:"primaryKey" example:"1"`
	CategoryID		uint	`json:"category_id"`
	QuestionText	string	`json:"question_text" gorm:"not null"`
	Answers			[]Answer	`json:"answers"`
}

type QuestionCreate struct {
	CategoryID		uint		`json:"category_id" gorm:"not null"`
	QuestionText	string		`json:"question_text" gorm:"not null"  example:"Text of question"`
	Answers			[]AnswerCreate	`json:"answers"`
}

type Answer struct {
	Id      		uint    `json:"id" gorm:"primaryKey" example:"1"`
	QuestionID		uint	`json:"question_id"`
	AnswerText		string	`json:"answer_text" example:"Text of answer item"`
	IsTrue			bool	`json:"is_true"`
}

type AnswerCreate struct {
	AnswerText		string	`json:"answer_text" example:"Text of answer item"`
	IsTrue			bool	`json:"is_true"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
}