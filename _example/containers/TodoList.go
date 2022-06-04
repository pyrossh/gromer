package containers

import (
	"context"

	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
`)

type TodoListProps struct {
	ID     string `json:"id"`
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

func TodoList(ctx context.Context, props TodoListProps) (*Template, error) {
	index := Default(props.Page, 1)
	todos, err := todos.GetAllTodo(ctx, todos.GetAllTodoParams{
		Filter: props.Filter,
		Limit:  index,
	})
	if err != nil {
		return nil, err
	}
	return Html(`
		<ul id="todo-list" class="todo-list">
			{{#each todos as |todo|}}
				{{#Todo todo=todo}}{{/Todo}}
			{{/each}}
		</ul>
	`).Prop("todos", todos), nil
}
