package repositories

import (
	"errors"
	"golang-todos/database"
	"golang-todos/models"
	"golang-todos/types/requests"
	"time"

	"gorm.io/gorm"
)

type TodoRepositoryInterface interface {
	TodoShow(id string) (*models.Todo, error)
	TodoIndex() ([]models.Todo, error)
	TodoCreate(request *requests.CreateTodoRequest) (*models.Todo, error)
	TodoUpdate(id string, request *requests.UpdateTodoRequest) (*models.Todo, error)
	TodoDelete(id string) error
}

type DefaultTodoRepository struct{}

func (r *DefaultTodoRepository) TodoIndex() ([]models.Todo, error) {
	var todos []models.Todo

	if err := database.Database.Find(&todos).Error; err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}

func (r *DefaultTodoRepository) TodoShow(id string) (*models.Todo, error) {
	var todo models.Todo

	if err := database.Database.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not_found")
		}

		return nil, err
	}

	return &todo, nil
}

func (r *DefaultTodoRepository) TodoCreate(request *requests.CreateTodoRequest) (*models.Todo, error) {
	todo := models.Todo{
		Task:        request.Task,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	err := database.Database.Create(&todo).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *DefaultTodoRepository) TodoUpdate(id string, request *requests.UpdateTodoRequest) (*models.Todo, error) {
	var todo models.Todo

	if err := database.Database.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not_found")
		}

		return nil, err
	}

	now := time.Now()

	if request.Task != "" {
		todo.Task = request.Task
	}

	if request.IsComplete {
		todo.CompletedAt = &now
	}

	if !request.IsComplete {
		todo.CompletedAt = nil
	}

	todo.UpdatedAt = now

	err := database.Database.Save(&todo).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *DefaultTodoRepository) TodoDelete(id string) error {
	var todo models.Todo

	if err := database.Database.Unscoped().First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	if todo.DeletedAt != nil {
		return errors.New("deleted")
	}

	if err := database.Database.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}

var TodoRepository TodoRepositoryInterface = &DefaultTodoRepository{}
