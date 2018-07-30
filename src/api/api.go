package api

import (
	"encoding/json"
	"net/http"
	"service"
)

type Api struct {
	Service service.IQuizService
}

func (api Api) GetAllQuizHandler(w http.ResponseWriter, r *http.Request) {
	quizzes, err := api.Service.GetQuizzes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(quizzes)
}
