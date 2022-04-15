package components

import (
	"github.com/pyros2097/gromer/_example/db"
	. "github.com/pyros2097/gromer/handlebars"
)

type TodoProps struct {
	Todo *db.Todo `json:"todo"`
}

func Todo(props TodoProps) *Template {
	return Html(`
		<tr>
			<td>{{ props.Todo.ID }}</td>
			<td>Book1</td>
			<td>Author1</td>
			<td>
					<button class="button is-primary">Edit</button>
			</td>
			<td>
					<button hx-swap="delete" class="button is-danger" hx-delete="/api/todos/{{ props.Todo.ID }}">Delete</button>
			</td>
		</tr>
	`)
}
