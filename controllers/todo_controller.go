package controllers

import (
	"net/http"

	"golang-todos/types"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}
