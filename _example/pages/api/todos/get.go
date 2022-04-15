package todos

import (
	"context"

	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/db"
)

type GetParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func GET(ctx context.Context, params GetParams) ([]*db.Todo, int, error) {
	limit := Default(params.Limit, 10)
	todos, err := db.Query.ListTodos(ctx, db.ListTodosParams{
		Limit:  int32(limit),
		Offset: int32(params.Offset),
	})
	if err != nil {
		return nil, 500, err
	}
	return todos, 200, nil
}
