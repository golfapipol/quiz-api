package service

import (
	"model"
)

func GetQuizzes() []model.Quiz {
	return make([]model.Quiz, 20)
}
