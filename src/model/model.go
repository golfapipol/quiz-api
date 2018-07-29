package model

import "gopkg.in/mgo.v2/bson"

type Quiz struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
}
