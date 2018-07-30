package service

import (
	"model"

	mgo "gopkg.in/mgo.v2"
)

type IQuizService interface {
	GetQuizzes() ([]model.Quiz, error)
}

type MockQuizService struct {
}

type QuizService struct {
	db *mgo.Session
}

func (mqs MockQuizService) GetQuizzes() ([]model.Quiz, error) {
	return make([]model.Quiz, 20), nil
}

func (qs QuizService) GetQuizzes() ([]model.Quiz, error) {
	var quizzes []model.Quiz
	err := qs.db.DB("quiz_api").C("quizzes").Find(nil).Limit(20).All(&quizzes)

	if err != nil {
		return []model.Quiz{}, err
	}
	return quizzes, nil
}

func GetQuizzes() []model.Quiz {
	return make([]model.Quiz, 20)
}
