package todos_todoId_

import (
	"context"

	"github.com/pyros2097/gromer/_example/services"
)

type PutParams struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func PUT(ctx context.Context, id string, params PutParams) (*services.Todo, int, error) {
	_, status, err := GET(ctx, id, GetParams{})
	if err != nil {
		return nil, status, err
	}
	todo, err := services.UpdateTodo(ctx, id, services.UpdateTodoParams{
		Text:      params.Text,
		Completed: params.Completed,
	})
	if err != nil {
		return nil, 500, err
	}
	return todo, 200, nil
}
