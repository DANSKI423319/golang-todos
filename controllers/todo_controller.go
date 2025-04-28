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
