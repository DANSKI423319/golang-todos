package controllers

import (
	"net/http"

	"golang-todos/repositories"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := repositories.TodoIndex()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := repositories.TodoShow(id)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
