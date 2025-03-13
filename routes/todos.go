package routes

import (
	"example/sample/controllers"

	"github.com/gin-gonic/gin"
)

func TodosRoutes(r *gin.Engine) {
	userGroup := r.Group("/todos")
	{
		userGroup.POST("/", controllers.TodosCreate)
		userGroup.GET("/", controllers.TodosIndex)
		userGroup.GET("/:id", controllers.TodosShow)
		userGroup.PUT("/:id", controllers.TodosUpdate)
		userGroup.DELETE("/:id", controllers.TodosDelete)
	}
}
