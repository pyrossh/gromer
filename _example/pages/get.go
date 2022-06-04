package pages

import (
	"context"

	_ "github.com/pyros2097/gromer/_example/components"
	. "github.com/pyros2097/gromer/handlebars"
)

type GetParams struct {
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

func GET(ctx context.Context, params GetParams) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="gromer example"}}
			<section class="todoapp">
				<header class="header">
					<h1>todos</h1>
					<form hx-post="/" hx-target="#todo-list" hx-swap="afterbegin" _="on htmx:afterOnLoad set #text.value to ''">
						<input type="hidden" name="intent" value="create" />
						<input class="new-todo" id="text" name="text" placeholder="What needs to be done?" autofocus="false" autocomplete="off">
					</form>	
				</header>	
				<section class="main">
					<input class="toggle-all" id="toggle-all" type="checkbox">
					<label for="toggle-all">Mark all as complete</label>
					{{#TodoList id="todo-list" page=page filter=filter }}{{/TodoList}}
				</section>
				<footer class="footer">
					{{#TodoCount page=page filter=filter}}{{/TodoCount}}
					<ul class="filters">
						<li>
							<a href="?filter=all">All</a>
						</li>
						<li>
							<a href="?filter=active">Active</a>
						</li>
						<li>
							<a href="?filter=completed">Completed</a>
						</li>
					</ul>
					<form hx-target="#todo-list" hx-post="/">
						<input type="hidden" name="intent" value="clear_completed" />
						<button type="submit" class="clear-completed" >Clear completed</button>
					</form>
				</footer>
			</section>
		{{/Page}}
		`).
		Prop("page", params.Page).
		Prop("filter", params.Filter).
		Render()
}
