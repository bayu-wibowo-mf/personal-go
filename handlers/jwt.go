package handlers

import (
	"example/sample/utils"

	"github.com/gin-gonic/gin"
)

func ProtectedHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized, token is required",
		})
		c.Abort()
		return
	}

	err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
