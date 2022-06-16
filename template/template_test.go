package template

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TodoData struct {
	ID        string
	Text      string
	Completed bool
}

func Todo(html Html, todo *TodoData) string {
	return html(`
		<li id={todo.ID} class={{ completed: todo.Completed }}>
			<div class="view">
				<span>{todo.Text}</span>
			</div>
		</li>
	`)
}

func WebsiteName() string {
	return "My Website"
}

func TestHtml(t *testing.T) {
	r := require.New(t)
	RegisterComponent(Todo, "todo")
	RegisterFunc(WebsiteName)
	ctx := map[string]interface{}{
		"_space": "",
		"todos": []*TodoData{
			{ID: "b1a7359c-ebb4-11ec-8ea0-0242ac120002", Text: "My first todo", Completed: true},
		},
	}
	actual := HtmlFunc(ctx)(`
		<ul id="todo-list" class="relative">
			<For key="todos" itemKey="todo">
				<Todo key="todo"></Todo>
			</For>
			<span>{WebsiteName}</span>
		</ul>
	`)
	expected := `<ul id="todo-list" class="relative">
</ul>`
	r.Equal(expected, actual)
}
