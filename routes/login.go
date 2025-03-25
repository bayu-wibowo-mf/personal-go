package routes

import (
	"example/sample/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine) {
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("/", controllers.Login)
	}
}
