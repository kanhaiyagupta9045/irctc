package controllers

import (
	"errors"
	"fmt"
	db "irctc/database"
	"irctc/models"
	"irctc/utils"
	"log"
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

		if err := c.BindJSON(&train); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}
		if err := validate.Struct(train); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}

		fmt.Println(train)

		if err := db.GetDB().Create(&train).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create train"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Train added successfully"})
	}
}

func GetSeatAvailability() gin.HandlerFunc {
	return func(c *gin.Context) {

		source := c.Query("src")
		destination := c.Query("dst")
		log.Println("Source:", source, "Destination:", destination)

		var trains []models.Train
		if err := db.GetDB().Where("source = ? AND destination = ?", source, destination).Find(&trains).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve trains does n't exist"})
			return
		}

		c.JSON(http.StatusOK, trains)
	}
}
