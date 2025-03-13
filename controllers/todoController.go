package controllers

import (
	"example/sample/initializers"
	"example/sample/models"

	"github.com/gin-gonic/gin"
)

func TodosCreate(c *gin.Context) {
	var body struct {
		Content string
		Status  bool
	}

	c.Bind(&body)

	todo := models.Todo{Content: body.Content, Status: body.Status}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosIndex(c *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func TodosShow(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	initializers.DB.First(&todo, id)

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Content string
		Status  bool
	}

	c.Bind(&body)

	var todo models.Todo
	initializers.DB.First(&todo, id)

	initializers.DB.Model(&todo).Updates(models.Todo{Content: body.Content, Status: body.Status})

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Todo{}, id)

	c.JSON(200, gin.H{
		"message": "Todo deleted",
	})
}
