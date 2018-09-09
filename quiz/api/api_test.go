package api_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	. "quiz/api"
	. "quiz/model"
	"quiz/router"
	"quiz/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

func Test_GetAllQuizHandler_Should_Be_20_Quiz(t *testing.T) {
	request := httptest.NewRequest("GET", "/v1/quizzes", nil)
	responseRecorder := httptest.NewRecorder()
	expectedQuizzes := 20
	api := ApiWithGin{
		Service: &service.MockQuizService{},
	}
	route := router.SetupRoute(&api)
	route.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var quizzes []Quiz
	json.Unmarshal(body, &quizzes)

	if expectedQuizzes != len(quizzes) {
		t.Errorf("Expected %d quiz but it got %d", expectedQuizzes, len(quizzes))
	}
}

func Test_CreateQuizHandler_Title_DISC_Description_Which_one_is_your_Should_Be_New_Quiz(t *testing.T) {
	quizRequest := QuizRequest{Title: "DISC", Description: "Which one is your?"}
	quizJSON, _ := json.Marshal(quizRequest)
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	expectedQuiz := Quiz{
		ID:          bson.NewObjectIdWithTime(fixedTime),
		Title:       "DISC",
		Description: "Which one is your?",
	}
	request := httptest.NewRequest("POST", "/v1/quizzes", bytes.NewBuffer(quizJSON))
	request.Header.Set("Content-Type", binding.MIMEJSON)
	responseRecorder := httptest.NewRecorder()

	api := ApiWithGin{
		Service: &service.MockQuizService{},
	}
	route := router.SetupRoute(&api)
	route.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var quiz Quiz
	json.Unmarshal(body, &quiz)

	if expectedQuiz != quiz {
		t.Errorf("Expected %v quiz but it got %v", expectedQuiz, quiz)
	}
}

func Test_UpdateQuizHandler_New_Description_DISC_is_type_of_people_Should_Be_Updated_Quiz(t *testing.T) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	oldQuiz := Quiz{
		ID:          bson.NewObjectIdWithTime(fixedTime),
		Title:       "DISC",
		Description: "Which one is your?",
	}
	updateQuizRequest := QuizRequest{
		Title:       "DISC",
		Description: "DISC is type of people",
	}
	quizJSON, _ := json.Marshal(updateQuizRequest)
	expectedQuiz := Quiz{
		ID:          bson.NewObjectIdWithTime(fixedTime),
		Title:       "DISC",
		Description: "DISC is type of people",
	}
	request := httptest.NewRequest("PUT", "/v1/quizzes/"+oldQuiz.ID.Hex(), bytes.NewBuffer(quizJSON))
	request.Header.Set("Content-Type", binding.MIMEJSON)
	responseRecorder := httptest.NewRecorder()

	api := ApiWithGin{
		Service: &service.MockQuizService{
			ExistedQuiz: oldQuiz,
		},
	}
	route := router.SetupRoute(&api)
	route.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var updatedQuiz Quiz
	json.Unmarshal(body, &updatedQuiz)

	if expectedQuiz != updatedQuiz {
		t.Errorf("Expected %v quiz but it got %v", expectedQuiz, updatedQuiz)
	}
}

func Test_GetQuizByIDHandler_Existed_ID_Should_Be_Quiz(t *testing.T) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	id := bson.NewObjectIdWithTime(fixedTime)
	expectedQuiz := Quiz{
		ID:          id,
		Title:       "DISC",
		Description: "DISC is type of people",
	}
	request := httptest.NewRequest("GET", "/v1/quizzes/"+id.Hex(), nil)
	request.Header.Set("Content-Type", binding.MIMEJSON)
	responseRecorder := httptest.NewRecorder()

	api := ApiWithGin{
		Service: &service.MockQuizService{
			ExistedQuiz: expectedQuiz,
		},
	}
	route := router.SetupRoute(&api)
	route.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var updatedQuiz Quiz
	json.Unmarshal(body, &updatedQuiz)

	if expectedQuiz != updatedQuiz {
		t.Errorf("Expected %v quiz but it got %v", expectedQuiz, updatedQuiz)
	}
}

func Test_DeleteQuizByIDHandler_Existed_ID_Should_Be_Status_OK(t *testing.T) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Jul-31")
	id := bson.NewObjectIdWithTime(fixedTime)
	expectedStatus := 200
	request := httptest.NewRequest("GET", "/v1/quizzes/"+id.Hex(), nil)
	request.Header.Set("Content-Type", binding.MIMEJSON)
	responseRecorder := httptest.NewRecorder()

	api := ApiWithGin{
		Service: &service.MockQuizService{},
	}
	route := router.SetupRoute(&api)
	route.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	if expectedStatus != response.StatusCode {
		t.Errorf("Expected status code: %v but it got %v", expectedStatus, response.StatusCode)
	}
}
