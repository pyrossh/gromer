package todos

import (
	"context"

	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services"
)

type GetParams struct {
	Limit int `json:"limit"`
}

func GET(ctx context.Context, params GetParams) ([]*services.Todo, int, error) {
	limit := Default(params.Limit, 10)
	todos := services.GetAllTodo(ctx, services.GetAllTodoParams{
		Limit: limit,
	})
	return todos, 200, nil
}
