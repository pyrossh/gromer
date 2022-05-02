package components

import (
	"github.com/pyros2097/gromer/_example/services"
	. "github.com/pyros2097/gromer/handlebars"
)

type TodoProps struct {
	Todo *services.Todo `json:"todo"`
}

func Todo(props TodoProps) *Template {
	return Html(`
		<li id="todo-{{ props.Todo.ID }}" {{#if props.Todo.Completed }} class="completed" {{/if}}>
			<div class="view">
				<input class="toggle" hx-post="/api/todos/{{ props.Todo.ID }}/complete" type="checkbox" {{#if props.Todo.Completed }} checked="" {{/if}} hx-target="#todo-{{ props.Todo.ID }}" hx-swap="outerHTML">
				<label hx-get="/todos/edit/{{ props.Todo.ID }}" hx-target="#todo-{{ props.Todo.ID }}" hx-swap="outerHTML">{{ props.Todo.Text }}</label>
				<button class="destroy" hx-delete="/api/todos/{{ props.Todo.ID }}" hx-target="#todo-{{ props.Todo.ID }}" hx-swap="delete"></button>
			</div>
		</li>
	`)
}
