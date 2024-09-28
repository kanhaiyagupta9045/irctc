package authentication

import (
	"errors"
	"fmt"
	db "irctc/database"
	"irctc/helpers"
	"irctc/models"
	"irctc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientToken := c.Request.Header.Get("Authorization")

		if clientToken == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authoriztion header Provided")})
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != nil {
			utils.ErrorResponse(c, err, http.StatusInternalServerError)
			return
		}
		var user models.User
		err = db.DB.Where("email = ?", claims.Email).First(&user).Error

		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, errors.New("user not found"), http.StatusNotFound)
			return
		} else if err != nil {
			utils.ErrorResponse(c, err, http.StatusInternalServerError)
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
