package todos_page

import (
	"context"

	. "github.com/pyros2097/gromer"
	_ "github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/pages/api/todos"
	. "github.com/pyros2097/gromer/handlebars"
)

type GetParams struct {
	Filter string `json:"filter"`
	Page   int    `json:"limit"`
}

func GET(ctx context.Context, params GetParams) (HtmlContent, int, error) {
	index := Default(params.Page, 1)
	todos, status, err := todos.GET(ctx, todos.GetParams{
		Filter: params.Filter,
		Limit:  index * 10,
	})
	if err != nil {
		return HtmlErr(status, err)
	}
	return Html(`
		{{#Page title="gromer example"}}
			{{#Header}}{{/Header}}
			<section class="todoapp">
				<header class="header">
					<h1>todos</h1>
					<form hx-post="/todos" hx-target="#todo-list" hx-swap="afterbegin" hx-ext="json-enc" _="on htmx:afterOnLoad set #text.value to ''">
						<input class="new-todo" id="text" name="text" placeholder="What needs to be done?" autofocus="false" autocomplete="off">
					</form>
				</header>
				<section class="main">
					<input class="toggle-all" id="toggle-all" type="checkbox">
					<label for="toggle-all">Mark all as complete</label>
					<ul id="todo-list" class="todo-list">
						{{#each todos as |todo|}}
							{{#Todo todo=todo}}{{/Todo}}
						{{/each}}
					</ul>
				</section>
				<footer class="footer">
					<!-- This should be '0 items left' by default-->
					<span class="todo-count" id="todo-count" hx-swap-oob="true">
						<strong>1</strong> item left </span>
					<!-- Remove this if you don't implement routing-->
					<ul class="filters">
						<li>
							<a href="/todos?filter=all">All</a>
						</li>
						<li>
							<a href="/todos?filter=active">Active</a>
						</li>
						<li>
							<a href="/todos?filter=completed">Completed</a>
						</li>
					</ul>
					<!-- Hidden if no completed items are left ↓-->
					<button class="clear-completed" hx-post="/todos/clear-completed" hx-target="#todo-list">Clear completed</button>
				</footer>
			</section>
		{{/Page}}
		`).
		Prop("todos", todos).
		Render()
}
