package containers

import (
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

var _ = Css(`
	.todo-count {
		float: left;
		text-align: left;
	}

	.todo-count strong {
		font-weight: 300;
	}
`)

func TodoCount(ctx Context, filter string) (*Node, error) {
	todos, err := todos.GetAllTodo(ctx, todos.GetAllTodoParams{
		Filter: filter,
		Limit:  1000,
	})
	if err != nil {
		return nil, err
	}
	ctx.Set("count", len(todos))
	return ctx.Render(`
		<span class="todo-count" id="todo-count" hx-swap-oob="true">
			<strong>{count}</strong> item left
		</span>
	`), nil
}
