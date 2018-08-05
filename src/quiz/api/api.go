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
	case http.MethodPut:
		api.UpdateQuizHandler(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}

func (api Api) GetAllQuizHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := api.Service.GetQuizzes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quizzes)
}

func (api Api) CreateQuizHandler(w http.ResponseWriter, r *http.Request) {
	var quiz model.Quiz
	json.NewDecoder(r.Body).Decode(&quiz)
	newQuiz, _ := api.Service.CreateQuiz(quiz)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newQuiz)
}

func (api Api) UpdateQuizHandler(w http.ResponseWriter, r *http.Request) {
	var quiz model.Quiz
	var id string
	json.NewDecoder(r.Body).Decode(&quiz)
	newQuiz, _ := api.Service.UpdateQuiz(id, quiz)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newQuiz)
}