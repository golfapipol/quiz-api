package api

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
)

type Api struct {
	Service service.IQuizService
}

func (api Api) QuizHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.GetAllQuizHandler(w, r)
	case http.MethodPost:
		api.CreateQuizHandler(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}

func (api Api) GetAllQuizHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := api.Service.GetQuizzes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(quizzes)
}

func (api Api) CreateQuizHandler(w http.ResponseWriter, r *http.Request) {
	var quiz model.Quiz
	json.NewDecoder(r.Body).Decode(&quiz)
	newQuiz, _ := api.Service.CreateQuiz(quiz)
	json.NewEncoder(w).Encode(newQuiz)
}
