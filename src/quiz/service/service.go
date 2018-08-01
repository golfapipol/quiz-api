package service

import (
	"errors"
	"model"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type IQuizService interface {
	GetQuizzes() ([]model.Quiz, error)
	CreateQuiz(quiz model.Quiz) (model.Quiz, error)
	UpdateQuiz(id string, quiz model.Quiz) (model.Quiz, error)
}

type QuizService struct {
	Session *mgo.Session
}

func (qs QuizService) GetQuizzes() ([]model.Quiz, error) {
	quizzes := make([]model.Quiz, 0)
	err := qs.Session.DB("quiz_api").C("quizzes").Find(nil).Limit(20).All(&quizzes)
	return quizzes, err
}

func (qs QuizService) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {
	quiz.ID = bson.NewObjectId()
	err := qs.Session.DB("quiz_api").C("quizzes").Insert(&quiz)
	return quiz, err
}

func (qs QuizService) UpdateQuiz(id string, quiz model.Quiz) (model.Quiz, error) {
	var existedQuiz model.Quiz
	err := qs.Session.DB("quiz_api").C("quizzes").FindId(id).One(&existedQuiz)
	if err != nil {
		return quiz, err
	}
	if existedQuiz.ID.Hex() != id {
		return quiz, errors.New("Not found")
	}
	change := bson.M{}
	if quiz.Title != "" {
		change["title"] = quiz.Title
	}
	if quiz.Description != "" {
		change["description"] = quiz.Description
	}
	err = qs.Session.DB("quiz_api").C("quizzes").UpdateId(id, change)
	if err != nil {
		return quiz, err
	}
	err = qs.Session.DB("quiz_api").C("quizzes").FindId(id).One(&existedQuiz)
	return existedQuiz, err
}
