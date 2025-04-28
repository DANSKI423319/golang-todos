package repositories

import (
	"golang-todos/database"
	"golang-todos/types"
)

func TodoIndex() ([]types.Todo, error) {
	var todos []types.Todo

	if err := database.Database.Find(&todos).Error; err != nil {
		return []types.Todo{}, err
	}

	return todos, nil
}
