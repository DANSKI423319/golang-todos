package routes

import (
	"golang-todos/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine) {
	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:id", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
}
