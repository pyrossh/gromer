package todos_todoId_

import (
	"context"
	"fmt"
	"strings"

	"github.com/pyros2097/gromer/example/db"
)

type GetParams struct {
	Show string `json:"show"`
}

func GET(ctx context.Context, id string, params GetParams) (*db.Todo, int, error) {
	todo, err := db.Query.GetTodo(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, 404, fmt.Errorf("Todo with id '%s' not found", id)
		}
		return nil, 500, err
	}
	if params.Show == "true" {
		todo.Completed = true
	}
	return todo, 200, nil
}
