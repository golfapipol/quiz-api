package api

import (
	"encoding/json"
	"net/http"
	"service"
)

func GetAllQuizHandler(w http.ResponseWriter, r *http.Request) {
	quizzes := service.GetQuizzes()
	json.NewEncoder(w).Encode(quizzes)
}
