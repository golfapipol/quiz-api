package main

import (
	"fmt"
	"os"
	apiLibrary "quiz/api"
	"quiz/config"
	"quiz/router"
	"quiz/service"

	mgo "gopkg.in/mgo.v2"
)

const url = "mongodb://localhost:27017"

func main() {
	environment := "development"
	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	}
	configs := config.GetConfiguration(environment)

	DBConnection, err := mgo.Dial(configs["mongo"])
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
