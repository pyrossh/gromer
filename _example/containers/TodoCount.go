package containers

import (
	"context"

	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/handlebars"
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

type TodoCountProps struct {
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

func TodoCount(ctx context.Context, props TodoCountProps) (*Template, error) {
	index := Default(props.Page, 1)
	todos, err := todos.GetAllTodo(ctx, todos.GetAllTodoParams{
		Filter: props.Filter,
		Limit:  index,
	})
	if err != nil {
		return nil, err
	}
	return Html(`
		<span class="todo-count" id="todo-count" hx-swap-oob="true">
			<strong>{{ count }}</strong> item left
		</span>
	`).Prop("count", len(todos)), nil
}
