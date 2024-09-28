package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, err error, code int) {
	c.JSON(code, gin.H{"error": err.Error()})
}

func SuccessResponse(c *gin.Context, data string, code int) {
	c.JSON(code, gin.H{"message": data})
}
