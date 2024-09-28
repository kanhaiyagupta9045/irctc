package routes

import (
	"irctc/authentication"
	"irctc/controllers"

	"github.com/gin-gonic/gin"
)

func TrainRoutes(router *gin.Engine) {
	router.GET("/seat-availability", controllers.GetSeatAvailability())
	router.Use(authentication.Authenticate())

	router.POST("/add/train", controllers.AddTrain())

}
