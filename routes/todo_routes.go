package routes

import (
	"golang-todos/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine) {
	r.GET("/todos", controllers.GetTodos)
}
