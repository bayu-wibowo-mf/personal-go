package main

import (
	"example/sample/initializers"
	"example/sample/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	routes.TodosRoutes(r)
	routes.EmployeeDemographyRoutes(r)

	r.Run()
}
