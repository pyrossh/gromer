package containers

import (
	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

func TodoList(c *Context, page int, filter string) []*Tag {
	// c.Styles(M{
	// 	"container": "list-none",
	// })
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
		<ul id="todo-list" class="TodoList">	
			for i, v := range todos {
				return (
					<Todo todo={v} />
				)
			}
		</ul>
	`)
}
