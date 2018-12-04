package model

import "github.com/globalsign/mgo/bson"

type QuizRequest struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type Quiz struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
}

type QuizResponse Quiz
