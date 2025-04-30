package routes

import (
	"golang-todos/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.RouterGroup) {
	todoRoutes := r.Group("/todos")
	{
		todoRoutes.GET("", controllers.GetTodos)
		todoRoutes.GET("/:id", controllers.GetTodo)
		todoRoutes.POST("", controllers.CreateTodo)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}
}
