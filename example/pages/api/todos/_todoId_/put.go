package todos_todoId_

import (
	"context"
	"time"

	"github.com/pyros2097/gromer/example/db"
)

type PutParams struct {
	Completed bool `json:"completed"`
}

func PUT(ctx context.Context, id string, params PutParams) (*db.Todo, int, error) {
	todo, err := db.Query.UpdateTodo(ctx, db.UpdateTodoParams{
		ID:        id,
		Completed: params.Completed,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, 500, err
	}
	return todo, 200, nil
}
