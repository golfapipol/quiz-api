package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	. "model"
	"net/http/httptest"
	"service"
	"strings"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func Test_GetAllQuizHandler_Should_Be_20_Quiz(t *testing.T) {
	request := httptest.NewRequest("GET", "/v1/quizzes", nil)
	responseRecorder := httptest.NewRecorder()
	expectedQuizzes := 20
	api := Api{
		Service: &service.MockQuizService{},
	}
	api.QuizHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var quizzes []Quiz
	json.Unmarshal(body, &quizzes)

	if expectedQuizzes != len(quizzes) {
		t.Errorf("Expected %d quiz but it got %d", expectedQuizzes, len(quizzes))
	}
}

func Test_CreateQuizHandler_Title_DISC_Description_Which_one_is_your_Should_Be_New_Quiz(t *testing.T) {
	quizRequest := Quiz{Title: "DISC", Description: "Which one is your?"}
	quizJSON, _ := json.Marshal(quizRequest)
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	expectedQuiz := Quiz{
		ID:          bson.NewObjectIdWithTime(fixedTime),
		Title:       "DISC",
		Description: "Which one is your?",
	}
	request := httptest.NewRequest("POST", "/v1/quizzes", strings.NewReader(string(quizJSON)))
	responseRecorder := httptest.NewRecorder()

	api := Api{
		Service: &service.MockQuizService{},
	}
	api.QuizHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var quiz Quiz
	json.Unmarshal(body, &quiz)

	if expectedQuiz != quiz {
		t.Errorf("Expected %v quiz but it got %v", expectedQuiz, quiz)
	}
}
