package todos

import (
	"context"

	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services"
)

type GetParams struct {
	Limit  int    `json:"limit"`
	Filter string `json:"filter"`
}

func GET(ctx context.Context, params GetParams) ([]*services.Todo, int, error) {
	limit := Default(params.Limit, 10)
	todos := services.GetAllTodo(ctx, services.GetAllTodoParams{
		Limit: limit,
	})
	if params.Filter == "completed" {
		newTodos := []*services.Todo{}
		for _, v := range todos {
			if v.Completed {
				newTodos = append(newTodos, v)
			}
		}
		return newTodos, 200, nil
	}
	if params.Filter == "active" {
		newTodos := []*services.Todo{}
		for _, v := range todos {
			if !v.Completed {
				newTodos = append(newTodos, v)
			}
		}
		return newTodos, 200, nil
	}
	return todos, 200, nil
}
