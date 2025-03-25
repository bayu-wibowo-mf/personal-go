package routes

import (
	"example/sample/controllers"
	"example/sample/handlers"

	"github.com/gin-gonic/gin"
)

func TodosRoutes(r *gin.Engine) {
	userGroup := r.Group("/todos", handlers.ProtectedHandler)
	{
		userGroup.POST("/", controllers.TodosCreate)
		userGroup.GET("/", controllers.TodosIndex)
		userGroup.GET("/:id", controllers.TodosShow)
		userGroup.PUT("/:id", controllers.TodosUpdate)
		userGroup.DELETE("/:id", controllers.TodosDelete)
	}
}
