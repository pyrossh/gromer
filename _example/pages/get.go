package pages

import (
	"context"

	. "github.com/pyros2097/gromer"
	_ "github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/pages/api/todos"
	. "github.com/pyros2097/gromer/handlebars"
)

type GetParams struct {
	Page int `json:"limit"`
}

func GET(ctx context.Context, params GetParams) (HtmlContent, int, error) {
	page := Default(params.Page, 1)
	todos, status, err := todos.GET(ctx, todos.GetParams{
		Limit: page * 10,
	})
	if err != nil {
		return HtmlErr(status, err)
	}
	return Html(`
		{{#Page title="gromer example"}}
			{{#Header}}{{/Header}}
			<main class="box center">
				<div class="columns">
					<div>
						<img src="/assets/icon.png" width="48" height="48" />
					</div>
					<div style="margin-left: 8px;">
						<h1>Hello this is a h1</h1>
						<h2>Hello this is a h2</h2>
					</div>
				</div>
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
								{{#Todo todo=todo}}{{/Todo}}
							{{/each}}
						</tbody>
				</table>
				<nav class="pagination" role="navigation" aria-label="pagination">
					<a class="pagination-previous">Previous</a>
					<a class="pagination-next">Next page</a>
					<ul class="pagination-list">
						<li>
							<a class="pagination-link" aria-label="Goto page 1">1</a>
						</li>
						<li>
							<a class="pagination-link" aria-label="Goto page 2">2</a>
						</li>
					</ul>
				</nav>
			</main>
		{{/Page}}
		`).
		Prop("todos", todos).
		Render()
}
