package todos_page

import (
	"context"

	"github.com/pyros2097/gromer/_example/pages/api/todos"
	. "github.com/pyros2097/gromer/handlebars"
)

func POST(ctx context.Context, b todos.PostParams) (HtmlContent, int, error) {
	todo, status, err := todos.POST(ctx, b)
	if err != nil {
		return HtmlErr(status, err)
	}
	return Html(`
		{{#Todo todo=todo}}{{/Todo}}
	`).
		Prop("todo", todo).
		Render()
}
