package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	. "model"
	"net/http/httptest"
	"service"
	"testing"
)

func Test_GetAllQuizHandler_Should_Be_20_Quiz(t *testing.T) {
	request := httptest.NewRequest("GET", "/v1/quizzes", nil)
	responseRecorder := httptest.NewRecorder()
	expectedQuizzes := 20
	api := Api{
		Service: &service.MockQuizService{},
	}
	api.GetAllQuizHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var quizzes []Quiz
	json.Unmarshal(body, &quizzes)

	if expectedQuizzes != len(quizzes) {
		t.Errorf("Expected %d quiz but it got %d", expectedQuizzes, len(quizzes))
	}
}
