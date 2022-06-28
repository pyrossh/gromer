package routes

import (
	_ "github.com/pyros2097/gromer/_example/components"
	. "github.com/pyros2097/gromer/gsx"
)

var (
	Meta = M{
		"title":       "Gromer Todos",
		"description": "Gromer Todos",
		"author":      "gromer",
		"keywords":    "gromer",
	}

	Styles = M{
		"container":       "bg-gray-50 min-h-screen font-sans",
		"todos-container": "container mx-auto flex flex-col items-center",
		"title":           "text-opacity-20 text-red-900 text-8xl text-center",
		"main":            "mt-8 shadow-xl w-full max-w-prose bg-white",
		"bottom": M{
			"container": "flex flex-row items-center flex-wrap sm:flex-nowrap p-2 font-light border-t-2 border-gray-100",
			"row":       "flex-1 flex flex-row",
			"section-1": "flex-1 flex flex-row order-1 justify-start",
			"section-2": "flex-1 flex flex-row order-2 sm:order-3 justify-end",
			"section-3": "flex-1 flex flex-row order-3 sm:order-2 min-w-full sm:min-w-min justify-center",
			"link":      "rounded border px-1 mx-2 hover:border-red-100",
			"active":    "border-red-900",
			"clear":     "font-light hover:underline",
			"disabled":  "invisible disabled",
		},
		"footer": M{
			"container": "mt-16 p-4 flex flex-col",
			"link":      "hover:underline",
			"subtitle":  "m-0.5 text-xs text-center text-gray-500",
		},
	}
)

type GetParams struct {
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

func GET(c *Context, params GetParams) (*Node, int, error) {
	return c.Render(`
		<div class="container">
			<div class="todos-container">
				<header>
					<h1 class="title">todos</h1>
					<form hx-post="/" hx-target="#todo-list" hx-swap="afterbegin" _="on htmx:afterOnLoad set #text.value to ''">
						<input type="hidden" name="intent" value="create" />
						<input class="new-todo" id="text" name="text" placeholder="What needs to be done?" autofocus="false" autocomplete="off" />
					</form>	
				</header>
				<main class="main">
					<input class="toggle-all" id="toggle-all" type="checkbox" />
					<label for="toggle-all">Mark all as complete</label>
					<TodoList id="todo-list" page="{params.Page}" filter="{params.Filter}"></TodoList>
					<div class="bottom">
						<div class="section-1">
							<TodoCount filter="{params.Filter}"></TodoCount>
						</div>
						<ul class="section-2">
							<li>
								<a href="?filter=all" class="link active">All</a>
							</li>
							<li>
								<a href="?filter=active" class="link">Active</a>
							</li>
							<li>
								<a href="?filter=completed" class="link">Completed</a>
							</li>
						</ul>
						<div class="section-3">
						<form hx-target="#todo-list" hx-post="/">
							<input type="hidden" name="intent" value="clear_completed" />
							<button type="submit" class="bottom-clear">Clear completed</button>
						</form>
						</div>
					</div>
				</main>
				<footer class="footer">
					<span class="subtitle">Written by 
						<a class="link" href="https://github.com/pyrossh/">pyrossh</a>
					</span>
					<span class="subtitle">using 
						<a class="link" href="https://github.com/pyrossh/gromer">Gromer</a>	
					</span>
					<span class="subtitle">thanks to 
						<a class="link" href="https://github.com/wishawa/">Wisha Wa</a>
					</span>
					<span class="subtitle">according to the spec 
						<a class="link" href="https://todomvc.com/">TodoMVC</a>
					</span>	
				</footer>
			</div>
		</div>
	`), 200, nil
}
