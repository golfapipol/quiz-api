package service_test

import (
	. "model"
	. "service"
	"testing"
)

func Test_GetQuizzes_Should_Be_20_Quizzes(t *testing.T) {
	expected := make([]Quiz, 20)
	actual := GetQuizzes()
	if len(expected) != len(actual) {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}
