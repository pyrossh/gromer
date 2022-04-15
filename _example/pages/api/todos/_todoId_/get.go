package todos_todoId_

import (
	"context"

	"github.com/pyros2097/gromer/_example/services"
)

type GetParams struct {
	Show string `json:"show"`
}

func GET(ctx context.Context, id string, params GetParams) (*services.Todo, int, error) {
	todo, err := services.GetTodo(ctx, id)
	if err != nil {
		return nil, 500, err
	}
	if params.Show == "true" {
		todo.Completed = true
	}
	return todo, 200, nil
}
