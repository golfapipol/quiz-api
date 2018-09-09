package router

import (
	apiLibrary "quiz/api"

	"github.com/gin-gonic/gin"
)

func SetupRoute(api apiLibrary.Api) *gin.Engine {
	router := gin.Default()
	router.GET("/v1/quizzes", api.GetAllQuizHandler)
	router.POST("/v1/quizzes", api.CreateQuizHandler)
	router.GET("/v1/quizzes/:id", api.GetQuizByIdHandler)
	router.PUT("/v1/quizzes/:id", api.UpdateQuizHandler)
	router.DELETE("/v1/quizzes/:id", api.DeleteQuizByIDHandler)
	return router
}
