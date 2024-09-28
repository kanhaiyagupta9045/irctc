package routes

import (
	"irctc/authentication"
	"irctc/controllers"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	router.Use(authentication.Authenticate())

	router.POST("/add/train", controllers.AddTrain())
}
