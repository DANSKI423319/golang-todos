package requests

type CreateTodoRequest struct {
	Task string `form:"task" json:"task" validate:"required,min=1,max=255"`
}

// Remove strict requirements here so that we can update the task and is_complete fields
// without needing to provide all fields at the same time
type UpdateTodoRequest struct {
	Task       string `form:"task" json:"task" validate:"max=255"`
	IsComplete bool   `form:"is_complete" json:"is_complete" validate:"boolean"`
}
