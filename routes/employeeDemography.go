package routes

import (
	"example/sample/controllers"
	"example/sample/handlers"

	"github.com/gin-gonic/gin"
)

func EmployeeDemographyRoutes(r *gin.Engine) {
	userGroup := r.Group("/employeeDemography", handlers.ProtectedHandler)
	{
		userGroup.GET("/:id", controllers.EmployeeDemographyShow)
		userGroup.GET("/", controllers.EmployeeDemographyIndex)
	}
}
