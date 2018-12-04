package main

import (
	"quiz/internal/service"
	"quiz/internal/api"
	"quiz/internal/router"
	"fmt"
	"log"
	"os"

	"github.com/globalsign/mgo"
)

func main() {
	DBConnection, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal("Cannot connect database ", err.Error())
	}
	quizService := service.MongoQuizService{Session: DBConnection}
	quizAPI := api.QuizAPI{Service: quizService}

	server := router.SetupRoute(quizAPI)
	server.Run(":3000")
	fmt.Println("Quiz Api is Listening")
}
