package api

import (
	"fmt"
	"net/http"
	"quiz/model"
	"quiz/service"

	"github.com/gin-gonic/gin"
)

type Api interface {
	GetAllQuizHandler(context *gin.Context)
	CreateQuizHandler(context *gin.Context)
	UpdateQuizHandler(context *gin.Context)
}

type ApiWithGin struct {
	Service service.QuizService
}

func (api ApiWithGin) GetAllQuizHandler(context *gin.Context) {
	quizzes, err := api.Service.GetQuizzes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, quizzes)
}

func (api *ApiWithGin) CreateQuizHandler(context *gin.Context) {
	var quiz model.QuizRequest
	if err := context.Bind(&quiz); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQuiz, _ := api.Service.CreateQuiz(quiz)
	context.JSON(http.StatusOK, newQuiz)
}

func (api *ApiWithGin) UpdateQuizHandler(context *gin.Context) {
	var quiz model.QuizRequest
	id := context.Param("id")
	fmt.Println("id", id)
	if err := context.Bind(&quiz); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("quiz", quiz)
	newQuiz, _ := api.Service.UpdateQuiz(id, quiz)
	fmt.Println("quiz", newQuiz)
	context.JSON(http.StatusOK, newQuiz)
}
