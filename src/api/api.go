package api

import (
	"encoding/json"
	"net/http"
	"service"
)

type Api struct {
	Service service.QuizService
}

func (api Api) GetAllQuizHandler(w http.ResponseWriter, r *http.Request) {
	quizzes := api.Service.GetQuizzes()
	json.NewEncoder(w).Encode(quizzes)
}
