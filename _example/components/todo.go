package components

import (
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
`)

type TodoProps struct {
	Todo *todos.Todo `json:"todo"`
}

func Todo(props TodoProps) *Template {
	return Html(`
		<li id="todo-{{ props.Todo.ID }}" {{#if props.Todo.Completed }} class="completed" {{/if}}>
			<div class="view">
				<form  hx-target="#todo-{{ props.Todo.ID }}" hx-swap="outerHTML">
					<input type="hidden" name="intent" value="complete" />
					<input type="hidden" name="id" value="{{ props.Todo.ID }}" />
					<input hx-post="/" class="checkbox" type="checkbox" {{#if props.Todo.Completed }} checked="" {{/if}} />
				</form>
				<label>{{ props.Todo.Text }}</label>
				<!-- <label hx-get="/todos/edit/{{ props.Todo.ID }}" hx-target="#todo-{{ props.Todo.ID }}" hx-swap="outerHTML">{{ props.Todo.Text }}</label> -->
				<form hx-post="/" hx-target="#todo-{{ props.Todo.ID }}" hx-swap="delete">
					<input type="hidden" name="intent" value="delete" />
					<input type="hidden" name="id" value="{{ props.Todo.ID }}" />
					<button class="destroy"></button>
				</form>
			</div>
		</li>
	`)
}
