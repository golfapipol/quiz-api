package main

import (
	"fmt"
	apiLibrary "quiz/api"
	"quiz/router"
	"quiz/service"

	mgo "gopkg.in/mgo.v2"
)

const url = "mongodb://localhost:27017"

func main() {

	DBConnection, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}
	quizService := service.MongoQuizService{Session: DBConnection}
	api := apiLibrary.ApiWithGin{Service: quizService}

	server := router.SetupRoute(&api)
	server.Run(":3000")
	fmt.Println("Quiz Api is Listening")
}
