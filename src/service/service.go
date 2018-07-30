package service

import (
	"model"
)

type QuizService interface {
	GetQuizzes() []model.Quiz
}

type MockQuizService struct {
}

func (mq MockQuizService) GetQuizzes() []model.Quiz {
	return make([]model.Quiz, 20)
}

func GetQuizzes() []model.Quiz {
	return make([]model.Quiz, 20)
}
