package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(c *gin.Context, err error) {
	validationErrors := err.(validator.ValidationErrors)
	firstError := validationErrors[0] // Get the first validation error
	c.JSON(http.StatusBadRequest, gin.H{
		"field":   firstError.Field(),
		"message": firstError.ActualTag(),
	})
}

func HandleAmbiguousError(c *gin.Context, err error) {
	if err.Error() == "not_found" {
		HandleItemNotFoundError(c)
		return
	}
	if err.Error() == "deleted" {
		HandleItemNotFoundError(c)
		return
	}
	HandleInternalServerError(c, err)
}

func HandleItemNotFoundError(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func HandleInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func HandleBindingError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
