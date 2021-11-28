package todos_todoId_

import (
	"context"

	"github.com/pyros2097/gromer/example/db"
)

type GetParams struct {
	Show string `json:"show"`
}

func GET(ctx context.Context, id string, params GetParams) (*db.Todo, int, error) {
	todo, err := db.Query.GetTodo(ctx, id)
	if err != nil {
		return nil, 500, err
	}
	if params.Show == "true" {
		todo.Completed = true
	}
	return todo, 200, nil
}
