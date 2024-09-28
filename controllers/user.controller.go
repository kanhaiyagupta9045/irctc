package controllers

import (
	"errors"
	db "irctc/database"
	"irctc/helpers"
	"irctc/models"
	"irctc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}
		if err := validate.Struct(user); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}

		var existingUser models.User
		err := db.GetDB().Where("email = ?", user.Email).First(&existingUser).Error
		if err == nil {
			utils.ErrorResponse(c, errors.New("email already exists"), http.StatusConflict)
			return
		}

		if err != nil && err != gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, err, http.StatusInternalServerError)
			return
		}

		hashedPassword, err := helpers.HashPassword(user.Password)
		if err != nil {
			utils.ErrorResponse(c, err, http.StatusInternalServerError)
		}
		user.Password = hashedPassword
		if err := db.GetDB().Create(&user).Error; err != nil {
			utils.ErrorResponse(c, err, http.StatusInternalServerError)
			return
		}
		utils.SuccessResponse(c, "User registered successfully", http.StatusCreated)
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var logindata models.LoginData
		if err := c.BindJSON(&logindata); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}
		if err := validate.Struct(logindata); err != nil {
			utils.ErrorResponse(c, err, http.StatusBadRequest)
			return
		}

		var existingUser models.User
		err := db.GetDB().Where("email = ?", logindata.Email).First(&existingUser).Error
		if err != nil || err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, errors.New("user does not exists"), http.StatusConflict)
			return
		}

		ok, msg := helpers.VerifyPassword(existingUser.Password, logindata.Password)

		if !ok {
			utils.ErrorResponse(c, errors.New(msg), http.StatusBadRequest)
			return
		}

		token, err := helpers.GenerateAccessToken(existingUser.Email)

		if err != nil {
			utils.ErrorResponse(c, errors.New("error while generating token"), http.StatusInternalServerError)
			return

		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}
