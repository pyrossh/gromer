package routes

import (
	"context"

	_ "github.com/pyros2097/gromer/_example/components"
	. "github.com/pyros2097/gromer/gsx"
)

var _ = Css(`
	.todoapp {
		background: #fff;
		margin: 130px 0 40px 0;
		position: relative;
		box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 25px 50px 0 rgba(0, 0, 0, 0.1);
	}

	.todoapp input::-webkit-input-placeholder {
		font-style: italic;
		font-weight: 300;
		color: #e6e6e6;
	}

	.todoapp input::-moz-placeholder {
		font-style: italic;
		font-weight: 300;
		color: #e6e6e6;
	}

	.todoapp input::input-placeholder {
		font-style: italic;
		font-weight: 300;
		color: #e6e6e6;
	}

	.todoapp h1 {
		position: absolute;
		top: -155px;
		width: 100%;
		font-size: 100px;
		font-weight: 100;
		text-align: center;
		color: rgba(175, 47, 47, 0.15);
		-webkit-text-rendering: optimizeLegibility;
		-moz-text-rendering: optimizeLegibility;
		text-rendering: optimizeLegibility;
	}
`)

// var (
// 	Container = Css(`
// 		background: #fff;
// 		margin: 130px 0 40px 0;
// 		position: relative;
// 		box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 25px 50px 0 rgba(0, 0, 0, 0.1);
// 	`)
// )

type GetParams struct {
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

// func IndexLoader() {
// 	Data: M{},
// 	Meta: []*Meta{},
// 	Links: []*Link{},
// }

func IndexAction(h Context, ctx context.Context, params PostParams) {
}

func IndexPage() {
}

// <meta name="description" content="{title}" />
// <meta name="author" content="pyrossh" />
// <meta name="keywords" content="pyros.sh, pyrossh, gromer" />
// <meta  />

func GET(h Context, params GetParams) (*Node, int, error) {
	h.Meta("title", "Todos")
	h.Set("page", params.Page)
	h.Set("filter", params.Filter)
	return h.Render(`
		<section class="todoapp">
			<header class="header">
				<h1>todos</h1>
				<form hx-post="/" hx-target="#todo-list" hx-swap="afterbegin" _="on htmx:afterOnLoad set #text.value to ''">
					<input type="hidden" name="intent" value="create" />
					<input class="new-todo" id="text" name="text" placeholder="What needs to be done?" autofocus="false" autocomplete="off" /s>
				</form>	
			</header>
			<section class="main">
				<input class="toggle-all" id="toggle-all" type="checkbox" />
				<label for="toggle-all">Mark all as complete</label>
				<TodoList id="todo-list" page="{page}" filter="{filter}" />
			</section>
			<footer class="footer">
				<TodoCount filter="{filter}" />
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
	`), 200, nil
}
