package routes

import (
	"irctc/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user/signup", controllers.SignUp())
	router.POST("/user/signin", controllers.LoginHandler())
}
