package service

import (
	"errors"
	"quiz/model"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

type QuizService interface {
	GetQuizzes() ([]model.Quiz, error)
	GetQuizByID(string) (model.Quiz, error)
	CreateQuiz(quiz model.QuizRequest) (model.Quiz, error)
	UpdateQuiz(id string, quiz model.QuizRequest) (model.Quiz, error)
	DeleteQuizByID(string) error
}

type MongoQuizService struct {
	Session *mgo.Session
}

func (qs MongoQuizService) GetQuizzes() ([]model.Quiz, error) {
	quizzes := make([]model.Quiz, 0)
	err := qs.Session.DB("quiz_api").C("quizzes").Find(nil).Limit(20).All(&quizzes)
	return quizzes, err
}

func (qs MongoQuizService) GetQuizByID(id string) (model.Quiz, error) {
	var existedQuiz model.Quiz
	err := qs.Session.DB("quiz_api").C("quizzes").FindId(bson.ObjectIdHex(id)).One(&existedQuiz)
	return existedQuiz, err
}

func (qs MongoQuizService) CreateQuiz(quiz model.QuizRequest) (model.Quiz, error) {
	var newQuiz = model.Quiz{}
	newQuiz.ID = bson.NewObjectId()
	newQuiz.Title = quiz.Title
	newQuiz.Description = quiz.Description
	err := qs.Session.DB("quiz_api").C("quizzes").Insert(&newQuiz)
	return newQuiz, err
}

func (qs MongoQuizService) UpdateQuiz(id string, quiz model.QuizRequest) (model.Quiz, error) {
	var existedQuiz model.Quiz
	err := qs.Session.DB("quiz_api").C("quizzes").FindId(bson.ObjectIdHex(id)).One(&existedQuiz)
	if err != nil {
		return existedQuiz, err
	}
	if existedQuiz.ID.Hex() != id {
		return existedQuiz, errors.New("Not found")
	}
	change := bson.M{}
	if quiz.Title != "" {
		change["title"] = quiz.Title
	}
	if quiz.Description != "" {
		change["description"] = quiz.Description
	}
	err = qs.Session.DB("quiz_api").C("quizzes").UpdateId(bson.ObjectIdHex(id), change)
	if err != nil {
		return existedQuiz, err
	}
	err = qs.Session.DB("quiz_api").C("quizzes").FindId(bson.ObjectIdHex(id)).One(&existedQuiz)
	return existedQuiz, err
}

func (qs MongoQuizService) DeleteQuizByID(id string) error {
	return qs.Session.DB("quiz_api").C("quizzes").Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
