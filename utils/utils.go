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
