package routes

import (
	"irctc/authentication"
	"irctc/controllers"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine) {
	router.Use(authentication.Authenticate())

	router.POST("/book/seat", controllers.BookSeat())
	router.GET("/booking/details", controllers.BookingDetails())
}
