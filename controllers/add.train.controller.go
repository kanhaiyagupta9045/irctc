package controllers

import (
	"errors"
	db "irctc/database"
	"irctc/models"
	"irctc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTrain() gin.HandlerFunc {
	return func(c *gin.Context) {
		var train models.Train
		user, exists := c.Get("user")
		if !exists {
			utils.ErrorResponse(c, errors.New("user not found"), http.StatusUnauthorized)
			return
		}
		authenticatedUser, ok := user.(models.User)
		if !ok {
			utils.ErrorResponse(c, errors.New("unable to cast user"), http.StatusInternalServerError)
			return
		}

		if authenticatedUser.UserType != ("ADMIN") {
			utils.ErrorResponse(c, errors.New("invalid user"), http.StatusBadRequest)
			return
		}
		if err := db.GetDB().Create(&train).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create train"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Train added successfully"})
	}
}
