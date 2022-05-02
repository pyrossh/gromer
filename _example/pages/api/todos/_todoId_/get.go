package todos_todoId_

import (
	"context"

	"github.com/pyros2097/gromer/_example/services"
)

func GET(ctx context.Context, id string) (*services.Todo, int, error) {
	todo, err := services.GetTodo(ctx, id)
	if err != nil {
		return nil, 500, err
	}
	return todo, 200, nil
}
