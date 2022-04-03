package todos_todoId_

import (
	"context"

	"github.com/pyros2097/gromer/example/db"
)

type DeleteParams struct{}

func DELETE(ctx context.Context, id string) (string, int, error) {
	_, status, err := GET(ctx, id, GetParams{})
	if err != nil {
		return "", status, err
	}
	err = db.Query.DeleteTodo(ctx, id)
	if err != nil {
		return id, 500, err
	}
	return id, 200, nil
}
