package requests

type CreateTodoRequest struct {
	Task string `form:"task" json:"task" validate:"required,min=1,max=255"`
}
