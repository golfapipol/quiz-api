package api

import (
	"net/http"
	"quiz/model"
	"quiz/service"

	"github.com/gin-gonic/gin"
)

type Api interface {
	GetAllQuizHandler(context *gin.Context)
	CreateQuizHandler(context *gin.Context)
	GetQuizByIdHandler(context *gin.Context)
	UpdateQuizHandler(context *gin.Context)
	DeleteQuizByIDHandler(context *gin.Context)
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
	if err := context.Bind(&quiz); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQuiz, err := api.Service.UpdateQuiz(id, quiz)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, newQuiz)
}

func (api ApiWithGin) GetQuizByIdHandler(context *gin.Context) {
	id := context.Param("id")
	newQuiz, err := api.Service.GetQuizByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, newQuiz)
}

func (api ApiWithGin) DeleteQuizByIDHandler(context *gin.Context) {
	id := context.Param("id")
	err := api.Service.DeleteQuizByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}
