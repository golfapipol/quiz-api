package api_test

import (
	"net/http/httptest"
	"testing"
)

func Test_GetAllQuizHandler_Should_Be_20_Quiz(t *testing.T) {
	request := httptest.NewRequest("GET", "/v1/quizzes", nil)
	responseRecorder := httptest.NewRecorder()
	expectedQuizzes := 20
	GetAllQuizHandler(responseRecorder, request)

	if expectedQuizzes != len(actual) {
		t.Errorf("Expected %d quiz but it got %d", expectedQuizzes, len(actual))
	}
}
