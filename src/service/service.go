package service

import (
	"model"
	"time"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type IQuizService interface {
	GetQuizzes() ([]model.Quiz, error)
	CreateQuiz(quiz model.Quiz) (model.Quiz, error)
}

type MockQuizService struct {
}

type QuizService struct {
	Session *mgo.Session
}

func (mqs MockQuizService) GetQuizzes() ([]model.Quiz, error) {
	return make([]model.Quiz, 20), nil
}

func (mqs MockQuizService) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	quiz.ID = bson.NewObjectIdWithTime(fixedTime)
	return quiz, nil
}

func (qs QuizService) GetQuizzes() ([]model.Quiz, error) {
	quizzes := make([]model.Quiz, 0)
	err := qs.Session.DB("quiz_api").C("quizzes").Find(nil).Limit(20).All(&quizzes)

	if err != nil {
		return []model.Quiz{}, err
	}
	return quizzes, nil
}

func (qs QuizService) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	quiz.ID = bson.NewObjectIdWithTime(fixedTime)
	return quiz, nil
}
