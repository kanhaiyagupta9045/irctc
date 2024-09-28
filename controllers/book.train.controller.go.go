package controllers

import (
	"errors"
	"fmt"
	db "irctc/database"
	"irctc/models"
	"irctc/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookSeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		var booking models.Booking
		if err := c.ShouldBindJSON(&booking); err != nil {
			utils.ErrorResponse(c, errors.New("invalid request body"), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(&booking); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}

		var train models.Train
		if err := db.GetDB().Where("train_number = ?", booking.TrainID).First(&train).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Train not found"})
			return
		}

		if train.AvailableSeats <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No available seats"})
			return
		}

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
		booking.UserID = uint32(authenticatedUser.ID)
		booking.CreatedAt = time.Now()
		booking.Status = "booked"
		booking.SeatNumber = train.TotalSeats - train.AvailableSeats + 1
		if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
			train.AvailableSeats--

			if err := tx.Save(&train).Error; err != nil {
				return err
			}

			if err := tx.Create(&booking).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book train"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Train booked successfully", "booking": booking})
	}
}

func BookingDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		userID := uint32(authenticatedUser.ID)
		fmt.Println(userID)
		var bookings []models.Booking
		if err := db.GetDB().Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
			utils.ErrorResponse(c, errors.New("could not retrieve bookings"), http.StatusInternalServerError)
			return
		}

		if len(bookings) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No bookings found for the user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"bookings": bookings})
	}
}
