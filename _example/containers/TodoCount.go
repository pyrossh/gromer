package containers

import (
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

func TodoCount(c *Context, filter string) *Node {
	todos, err := todos.GetAllTodo(c, todos.GetAllTodoParams{
		Filter: filter,
		Limit:  1000,
	})
	if err != nil {
		return Error(c, err)
	}
	count := 0
	for _, t := range todos {
		if !t.Completed {
			count++
		}
	}
	c.Set("count", count)
	return c.Render(`
		<span id="todo-count" class="todo-count" hx-swap-oob="true">
			<strong>{count}</strong> item left
		</span>
	`)
}
