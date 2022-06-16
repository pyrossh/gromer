package components

import (
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

var _ = Css(`
`)

func Todo(h Html, todo *todos.Todo) string {
	return h.Render(`
		<li id="todo-{todo.ID}" class={{ completed: todo.Completed }}>
			<div class="view">
				<form  hx-target="#todo-{todo.ID}" hx-swap="outerHTML">
					<input type="hidden" name="intent" value="complete" />
					<input type="hidden" name="id" value="{todo.ID}" />
					<input hx-post="/" class="checkbox" type="checkbox" checked={{ completed: todo.Completed }} />
				</form>
				<label>{todo.Text}</label>
				<form hx-post="/" hx-target="#todo-{todo.ID}" hx-swap="delete">
					<input type="hidden" name="intent" value="delete" />
					<input type="hidden" name="id" value="{todo.ID}" />
					<button class="destroy"></button>
				</form>
			</div>
		</li>
	`)
}

// <!-- <label hx-get="/todos/edit/{todo.ID}" hx-target="#todo-{todo.ID}" hx-swap="outerHTML">{{ props.Todo.Text }}</label> -->
