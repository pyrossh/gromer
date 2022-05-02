package todos_complete

import (
	"context"

	todos_todoId_ "github.com/pyros2097/gromer/_example/pages/api/todos/_todoId_"
	"github.com/pyros2097/gromer/_example/services"
)

func POST(ctx context.Context, id string) (*services.Todo, int, error) {
	_, status, err := todos_todoId_.GET(ctx, id)
	if err != nil {
		return nil, status, err
	}
	todo, err := services.UpdateTodo(ctx, id, services.UpdateTodoParams{
		Completed: true,
	})
	if err != nil {
		return nil, 500, err
	}
	return todo, 200, nil
}
