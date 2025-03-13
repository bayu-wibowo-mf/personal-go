package routes

import (
	"example/sample/controllers"

	"github.com/gin-gonic/gin"
)

func EmployeeDemographyRoutes(r *gin.Engine) {
	userGroup := r.Group("/employeeDemography")
	{
		userGroup.GET("/:id", controllers.EmployeeDemographyShow)
		userGroup.GET("/", controllers.EmployeeDemographyIndex)
	}
}
