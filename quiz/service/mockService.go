package service

import (
	"quiz/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MockQuizService struct {
	ExistedQuiz model.Quiz
}

func (mqs MockQuizService) GetQuizzes() ([]model.Quiz, error) {
	return make([]model.Quiz, 20), nil
}

func (mqs MockQuizService) GetQuizByID(id string) (model.Quiz, error) {
	return mqs.ExistedQuiz, nil
}

func (mqs MockQuizService) CreateQuiz(quiz model.QuizRequest) (model.Quiz, error) {
	var newQuiz = model.Quiz{}
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	newQuiz.ID = bson.NewObjectIdWithTime(fixedTime)
	newQuiz.Title = quiz.Title
	newQuiz.Description = quiz.Description
	return newQuiz, nil
}

func (mqs *MockQuizService) UpdateQuiz(id string, quiz model.QuizRequest) (model.Quiz, error) {
	if quiz.Title != "" {
		mqs.ExistedQuiz.Title = quiz.Title
	}
	if quiz.Description != "" {
		mqs.ExistedQuiz.Description = quiz.Description
	}
	return mqs.ExistedQuiz, nil
}

func (mqs MockQuizService) DeleteQuizByID(string) error {
	return nil
}
