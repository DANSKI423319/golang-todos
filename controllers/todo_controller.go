package controllers

import (
	"net/http"

	"golang-todos/types"

	"github.com/gin-gonic/gin"
)

// Mock data
var todos = []types.Todo{
	{ID: "1", Task: "Learn Go"},
	{ID: "2", Task: "Build an API"},
	{ID: "3", Task: "Dockerize the API"},
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}
