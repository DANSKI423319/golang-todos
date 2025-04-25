package main

import (
	"golang-todos/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine = gin.Default()

	routes.RegisterTodoRoutes(r)

	r.Run(":8080")
}
