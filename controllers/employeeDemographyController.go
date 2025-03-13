package controllers

import (
	"example/sample/initializers"
	"example/sample/models"

	"github.com/gin-gonic/gin"
)

func EmployeeDemographyShow(c *gin.Context) {
	id := c.Param("id")

	var employeeDemography models.EmployeeDemography
	initializers.DB.First(&employeeDemography, id)

	c.JSON(200, gin.H{
		"employeeDemography": employeeDemography,
	})
}

func EmployeeDemographyIndex(c *gin.Context) {
	var employeeDemography []models.EmployeeDemography
	initializers.DB.Find(&employeeDemography)

	c.JSON(200, gin.H{
		"employeeDemography": employeeDemography,
	})
}
