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
	todos, err := repositories.TodoRepository.TodoIndex()

	if err != nil {
		utils.HandleInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := repositories.TodoRepository.TodoShow(id)

	if err != nil {
		utils.HandleAmbiguousError(c, err)
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var request requests.CreateTodoRequest

	if err := c.ShouldBind(&request); err != nil {
		utils.HandleBindingError(c, err)
		return
	}

	if err := validate.Struct(request); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	todo, err := repositories.TodoRepository.TodoCreate(&request)
	if err != nil {
		utils.HandleInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	var request requests.UpdateTodoRequest

	if err := c.ShouldBind(&request); err != nil {
		utils.HandleBindingError(c, err)
		return
	}

	if err := validate.Struct(request); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	id := c.Param("id")
	todo, err := repositories.TodoRepository.TodoUpdate(id, &request)
	if err != nil {
		utils.HandleAmbiguousError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := repositories.TodoRepository.TodoDelete(id); err != nil {
		utils.HandleInternalServerError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
