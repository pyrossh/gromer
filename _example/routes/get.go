package routes

import (
	_ "github.com/pyros2097/gromer/_example/components"
	. "github.com/pyros2097/gromer/gsx"
)

var _ = Css(`
	.container {
		background: #fff;
		margin: 130px 0 40px 0;
		position: relative;
		box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 25px 50px 0 rgba(0, 0, 0, 0.1);
	}

	.container h1 {
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

	input::-webkit-input-placeholder {
		font-style: italic;
		font-weight: 300;
		color: #e6e6e6;
	}

	input::-moz-placeholder {
		font-style: italic;
		font-weight: 300;
		color: #e6e6e6;
	}

	input::input-placeholder {
		font-style: italic;
		font-weight: 300;
		color: #e6e6e6;
	}

	.clear-completed, .clear-completed:active {
		float: right;
		position: relative;
		line-height: 20px;
		text-decoration: none;
		cursor: pointer;
	}

	.clear-completed:hover {
		text-decoration: underline;
	}

	.filters {
		margin: 0;
		padding: 0;
		list-style: none;
		position: absolute;
		right: 0;
		left: 0;
	}

	.filters li {
		display: inline;
	}

	.filters li a {
		color: inherit;
		margin: 3px;
		padding: 3px 7px;
		text-decoration: none;
		border: 1px solid transparent;
		border-radius: 3px;
	}

	.filters li a:hover {
		border-color: rgba(175, 47, 47, 0.1);
	}

	.filters li a.selected {
		border-color: rgba(175, 47, 47, 0.2);
	}
`)

type GetParams struct {
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

func GET(c Context, params GetParams) (*Node, int, error) {
	c.Meta("title", "Gromer Todos")
	c.Meta("description", "Gromer Todos")
	c.Meta("author", "gromer")
	c.Meta("keywords", "gromer")
	return c.Render(`
		<div class="container">
			<header class="header">
				<h1>todos</h1>
				<form hx-post="/" hx-target="#todo-list" hx-swap="afterbegin" _="on htmx:afterOnLoad set #text.value to ''">
					<input type="hidden" name="intent" value="create" />
					<input class="new-todo" id="text" name="text" placeholder="What needs to be done?" autofocus="false" autocomplete="off" />
				</form>	
			</header>
			<section class="main">
				<input class="toggle-all" id="toggle-all" type="checkbox" />
				<label for="toggle-all">Mark all as complete</label>
				<TodoList id="todo-list" page="{params.Page}" filter="{params.Filter}" />
			</section>
			<footer class="footer">
				<TodoCount filter="{params.Filter}" />
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
		</div>
	`), 200, nil
}
