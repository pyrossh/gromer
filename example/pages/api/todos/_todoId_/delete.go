package todos_todoId_

import (
	"context"

	"github.com/pyros2097/gromer/example/db"
)

func DELETE(ctx context.Context, id string) (string, int, error) {
	err := db.Query.DeleteTodo(ctx, id)
	if err != nil {
		return id, 500, err
	}
	return id, 200, nil
}
