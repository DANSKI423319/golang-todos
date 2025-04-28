package repositories

import (
	"errors"
	"golang-todos/database"
	"golang-todos/models"

	"gorm.io/gorm"
)

func TodoIndex() ([]models.Todo, error) {
	var todos []models.Todo

	if err := database.Database.Find(&todos).Error; err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func TodoShow(id string) (models.Todo, error) {
	var todo models.Todo

	if err := database.Database.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Todo{}, nil
		}

		return models.Todo{}, err
	}

	return todo, nil
}
