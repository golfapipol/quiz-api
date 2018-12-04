package service_test

import (
	. "quiz/internal/model"
	. "quiz/internal/service"
	"testing"
	"time"

	bson "github.com/globalsign/mgo/bson"
)

func Test_GetQuizzes_Should_Be_20_Quizzes(t *testing.T) {
	expected := make([]Quiz, 20)
	service := MockQuizService{}
	actual, _ := service.GetQuizzes()
	if len(expected) != len(actual) {
		t.Errorf("expected %v but it got %v", len(expected), len(actual))
	}
}

func Test_CreateQuiz_Should_Be_New_Quiz(t *testing.T) {
	quiz := QuizRequest{Title: "DISC", Description: "Which one is your?"}
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	expectedQuiz := Quiz{
		ID:          bson.NewObjectIdWithTime(fixedTime),
		Title:       "DISC",
		Description: "Which one is your?",
	}
	service := MockQuizService{}
	actual, _ := service.CreateQuiz(quiz)
	if expectedQuiz != actual {
		t.Errorf("expected %v but it got %v", expectedQuiz, actual)
	}
}
