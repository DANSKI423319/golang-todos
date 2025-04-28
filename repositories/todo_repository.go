package repositories

import (
	"golang-todos/database"
	"golang-todos/models"
)

func TodoIndex() ([]models.Todo, error) {
	var todos []models.Todo

	if err := database.Database.Find(&todos).Error; err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}
