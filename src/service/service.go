package service

import (
	"model"
)

type QuizService interface {
	GetQuizzes() []model.Quiz
}

func GetQuizzes() []model.Quiz {
	return make([]model.Quiz, 20)
}
