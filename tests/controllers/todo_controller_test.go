package controllers_test

import (
	"encoding/json"
	"golang-todos/models"
	"golang-todos/repositories"
	"golang-todos/routes"
	"golang-todos/types/requests"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockTodoRepository struct{}

func TestTodoController(t *testing.T) {
	router := routes.SetupRouter()
	repositories.TodoRepository = &MockTodoRepository{}

	t.Run("todo index", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/todos", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Index Test")
	})

	t.Run("todo show", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/todos/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Test")
	})

	t.Run("todo create", func(t *testing.T) {
		payload := requests.CreateTodoRequest{
			Task: "Create Test",
		}
		payloadJSON, _ := json.Marshal(payload)
		println(string(payloadJSON))
		req := httptest.NewRequest("POST", "/api/v1/todos", strings.NewReader(string(payloadJSON)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Create Test")
	})

	t.Run("todo update", func(t *testing.T) {
		payload := requests.UpdateTodoRequest{
			Task: "Update Test",
		}
		payloadJSON, _ := json.Marshal(payload)
		req := httptest.NewRequest("PUT", "/api/v1/todos/1", strings.NewReader(string(payloadJSON)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Update Test")
	})

	t.Run("delete todo", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/api/v1/todos/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

// -- mocks

func (m *MockTodoRepository) TodoShow(id string) (models.Todo, error) {
	return models.Todo{
		ID:          1,
		Task:        "Show Test",
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTodoRepository) TodoIndex() ([]models.Todo, error) {
	return []models.Todo{
		{
			ID:          1,
			Task:        "Index Test",
			CompletedAt: nil,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (m *MockTodoRepository) TodoCreate(request *requests.CreateTodoRequest) (models.Todo, error) {
	todo := models.Todo{
		Task:        "Create Test",
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return todo, nil
}

func (m *MockTodoRepository) TodoUpdate(id string, request *requests.UpdateTodoRequest) (models.Todo, error) {
	todo := models.Todo{
		ID:          1,
		Task:        "Update Test",
		CompletedAt: nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return todo, nil
}

func (m *MockTodoRepository) TodoDelete(id string) error {
	return nil
}
