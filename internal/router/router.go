package router

import (
	"quiz/internal/api"

	"github.com/gin-gonic/gin"
)

func SetupRoute(quizAPI api.QuizAPI) *gin.Engine {
	router := gin.Default()
	router.GET("/v1/quizzes", quizAPI.GetAllQuizHandler)
	router.POST("/v1/quizzes", quizAPI.CreateQuizHandler)
	router.GET("/v1/quizzes/:id", quizAPI.GetQuizByIdHandler)
	router.PUT("/v1/quizzes/:id", quizAPI.UpdateQuizHandler)
	router.DELETE("/v1/quizzes/:id", quizAPI.DeleteQuizByIDHandler)
	return router
}
