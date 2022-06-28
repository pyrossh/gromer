package containers

import (
	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

var TodoListStyles = M{
	"container": "list-none",
}

func TodoList(c *Context, page int, filter string) *Node {
	index := Default(page, 1)
	todos, err := todos.GetAllTodo(c, todos.GetAllTodoParams{
		Filter: filter,
		Limit:  index,
	})
	if err != nil {
		return Error(c, err)
	}
	c.Set("todos", todos)
	return c.Render(`
		<ul id="todo-list" class="todolist" x-for="todo in todos">
			<Todo />
		</ul>
	`)
}
