package pages

import (
	"context"

	. "github.com/pyros2097/gromer"
	_ "github.com/pyros2097/gromer/example/components"
	"github.com/pyros2097/gromer/example/pages/api/todos"
)

type GetParams struct {
	Page int `json:"limit"`
}

func GET(ctx context.Context, params GetParams) (HtmlContent, int, error) {
	page := Default(params.Page, 1)
	todos, status, err := todos.GET(ctx, todos.GetParams{
		Limit:  10,
		Offset: 10 * (page - 1),
	})
	if err != nil {
		return HtmlErr(status, err)
	}
	return Html(`
		{{#Page "gromer example"}}
			<div class="flex flex-col justify-center items-center">
					{{#Header "123"}}
						A new link is here
					{{/Header}}
					<h1>Hello this is a h1</h1>
					<h2>Hello this is a h2</h2>
					<img src="/assets/icon.png" width="48" height="48" />
					<h3 x-data="{ message: 'I ❤️ Alpine' }" x-text="message">I ❤️ Alpine</h3>
					<table class="table">
							<thead>
									<tr>
										<th>ID</th>
										<th>Title</th>
										<th>Author</th>
									</tr>
							</thead>
							<tbody id="new-book" hx-target="closest tr" hx-swap="outerHTML swap:0.5s">
								{{#each todos as |todo|}}
									<tr>
										<td>{{todo.ID}}</td>
										<td>Book1</td>
										<td>Author1</td>
										<td>
												<button class="button is-primary">Edit</button>
										</td>
										<td>
												<button hx-swap="delete" class="button is-danger" hx-delete="/api/todos/{{todo.ID}}">Delete</button>
										</td>
									</tr>
								{{/each}}
							</tbody>
					</table>
			</div>
		{{/Page}}
		`, M{"todos": todos})
}
