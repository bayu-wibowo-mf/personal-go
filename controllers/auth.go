package controllers

import (
	"example/sample/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	if body.Username == "admin" && body.Password == "admin" {
		token, err := utils.GenerateToken(body.Username)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to generate token",
			})
		}
		c.JSON(200, gin.H{
			"message": "Login successful",
			"token":   token,
		})
	} else {
		c.JSON(401, gin.H{
			"message": "Invalid credentials",
		})
	}
}
