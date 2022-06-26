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

func TodoCount(c Context, filter string) (*Node, error) {
	todos, err := todos.GetAllTodo(c, todos.GetAllTodoParams{
		Filter: filter,
		Limit:  1000,
	})
	if err != nil {
		return nil, err
	}
	c.Set("count", len(todos))
	return c.Render(`
		<span id="todo-count" class="todo-count" hx-swap-oob="true">
			<strong>{count}</strong> item left
		</span>
	`), nil
}
