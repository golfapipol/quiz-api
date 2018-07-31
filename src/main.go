package main

import (
	apiLibrary "api"
	"fmt"
	"net/http"
	"service"

	mgo "gopkg.in/mgo.v2"
)

const url = "mongodb://localhost:27017"

func main() {

	DBConnection, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}
	quizService := service.QuizService{Session: DBConnection}
	api := apiLibrary.Api{Service: quizService}

	http.HandleFunc("/v1/quizzes", api.GetAllQuizHandler)
	http.ListenAndServe(":3000", nil)
	fmt.Println("Quiz Api is Listening")
}
