package controllers

import (
	"net/http"

	"golang-todos/repositories"
	"golang-todos/types/requests"
	"golang-todos/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

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

func CreateTodo(c *gin.Context) {
	var request requests.CreateTodoRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(request); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	todo, err := repositories.TodoCreate(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	var request requests.UpdateTodoRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(request); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	id := c.Param("id")
	todo, err := repositories.TodoUpdate(id, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := repositories.TodoDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
