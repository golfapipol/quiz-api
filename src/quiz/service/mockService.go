package service

import (
	"model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MockQuizService struct {
	existedQuiz model.Quiz
}

func (mqs MockQuizService) GetQuizzes() ([]model.Quiz, error) {
	return make([]model.Quiz, 20), nil
}

func (mqs MockQuizService) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	quiz.ID = bson.NewObjectIdWithTime(fixedTime)
	return quiz, nil
}

func (mqs *MockQuizService) UpdateQuiz(id string, quiz model.Quiz) (model.Quiz, error) {
	if quiz.Title != "" {
		mqs.existedQuiz.Title = quiz.Title
	}
	if quiz.Description != "" {
		mqs.existedQuiz.Description = quiz.Description
	}
	return mqs.existedQuiz, nil
}
