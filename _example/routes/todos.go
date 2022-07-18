package routes

import (
	_ "github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
	"github.com/rotisserie/eris"
)

var TodoStyles = M{
	"bg":        "bg-gray-50 min-h-screen font-sans",
	"container": "container mx-auto flex flex-col items-center",
	"title":     "text-opacity-20 text-red-900 text-8xl text-center",
	"main": M{
		"container":  "mt-8 shadow-xl w-full max-w-prose bg-white",
		"input-box":  "flex flex-row text-2xl h-16",
		"button":     "ml-4 w-8 disabled",
		"input-form": "flex flex-1",
		"input":      "flex-1 min-w-0 p-2 placeholder:text-gray-300",
	},
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

type TodosPageParams struct {
	Page   int    `json:"page"`
	Filter string `json:"filter"`
}

func getActive(v bool) string {
	if v {
		return "active"
	}
	return ""
}

func TodosPage(c *Context, params TodosPageParams) ([]*Tag, int, error) {
	allClass := getActive(params.Filter == "all")
	activeClass := getActive(params.Filter == "active")
	completedClass := getActive(params.Filter == "completed")
	c.Set("allClass", allClass)
	c.Set("activeClass", activeClass)
	c.Set("completedClass", completedClass)
	c.Meta(M{
		"title":       "Gromer Todos",
		"description": "Gromer Todos",
		"author":      "gromer",
		"keywords":    "gromer",
	})
	c.Styles(TodoStyles)
	return c.Render(`
		<div id="bg" class="bg">
			<div class="container">
				<header>
					<h1 class="title">"todos"</h1>
				</header>
				<main class="main">
					<div class="input-box">
						<form hx-target="#todo-list" hx-post="/">
							<input type="hidden" name="intent" value="select_all" />
							<button id="check-all" class="button" hx-swap-oob="true">
								<img src="/icons/check-all.svg?fill=gray-400" />
							</button>
						</form>
						<form class="input-form" hx-post="/" hx-target="#todo-list" hx-swap="afterbegin" _="on htmx:afterOnLoad set #text.value to ''">
							<input type="hidden" name="intent" value="create" />
							<input id="text" name="text" class="input" placeholder="What needs to be done?" autocomplete="off" />
						</form>
					</div>
					<TodoList id="todo-list" page={params.Page} filter={params.Filter} />
					<div class="bottom">
						<div class="section-1">
							<TodoCount filter={params.Filter} />
						</div>
						<ul class="section-2" hx-boost="true">
							<li>
								<a href="?filter=all" class="link {allClass}">"All"</a>
							</li>
							<li>
								<a href="?filter=active" class="link {activeClass}">"Active"</a>
							</li>
							<li>
								<a href="?filter=completed" class="link {completedClass}">"Completed"</a>
							</li>
						</ul>
						<div class="section-3">
						<form hx-target="#todo-list" hx-post="/">
							<input type="hidden" name="intent" value="clear_completed" />
							<button type="submit" class="bottom-clear">"Clear completed"</button>
						</form>
						</div>
					</div>
				</main>
				<div id="error">
				</div>
				<footer class="footer">
					<span class="subtitle">"Written by "
						<a class="link" href="https://github.com/pyrossh/">"pyrossh"</a>
					</span>
					<span class="subtitle">"using " 
						<a class="link" href="https://github.com/pyrossh/gromer">"Gromer"</a>	
					</span>
					<span class="subtitle">"thanks to" 
						<a class="link" href="https://github.com/wishawa/">"Wisha Wa"</a>
					</span>
					<span class="subtitle">"according to the spec "
						<a class="link" href="https://todomvc.com/">"TodoMVC"</a>
					</span>	
				</footer>
			</div>
		</div>
	`), 200, nil
}

type TodosActionParams struct {
	Intent string `json:"intent"`
	ID     string `json:"id"`
	Text   string `json:"text"`
}

func TodosAction(c *Context, params TodosActionParams) ([]*Tag, int, error) {
	if params.Intent == "select_all" {
		allTodos, err := todos.GetAllTodo(c, todos.GetAllTodoParams{
			Filter: "all",
			Limit:  1000,
		})
		if err != nil {
			return nil, 500, err
		}
		for _, t := range allTodos {
			_, err := todos.UpdateTodo(c, t.ID, todos.UpdateTodoParams{
				Text:      t.Text,
				Completed: true,
			})
			if err != nil {
				return nil, 500, err
			}
		}
		return c.Render(`
			<TodoCount filter="all" page="1" />
			<button id="check-all" class="button" hx-swap-oob="true">
				<img src="/icons/check-all.svg?fill=green-500" />
			</button>
			<TodoList id="todo-list" filter="all" page="1" />
		`), 200, nil
	} else if params.Intent == "clear_completed" {
		allTodos, err := todos.GetAllTodo(c, todos.GetAllTodoParams{
			Filter: "all",
			Limit:  1000,
		})
		if err != nil {
			return nil, 500, err
		}
		for _, t := range allTodos {
			if t.Completed {
				_, err := todos.DeleteTodo(c, t.ID)
				if err != nil {
					return nil, 500, err
				}
			}
		}
		return c.Render(`
			<TodoCount filter="all" page="1" />
			<TodoList id="todo-list" filter="all" page="1" />
		`), 200, nil
	} else if params.Intent == "create" {
		todo, err := todos.CreateTodo(c, params.Text)
		if err != nil {
			return nil, 500, err
		}
		c.Set("todo", todo)
		return c.Render(`
			<TodoCount filter="all" page="1" />
			<Todo />
		`), 200, nil
	} else if params.Intent == "delete" {
		_, err := todos.DeleteTodo(c, params.ID)
		if err != nil {
			return nil, 500, err
		}
		return nil, 200, nil
	} else if params.Intent == "complete" {
		todo, err := todos.GetTodo(c, params.ID)
		if err != nil {
			return nil, 500, err
		}
		_, err = todos.UpdateTodo(c, params.ID, todos.UpdateTodoParams{
			Text:      todo.Text,
			Completed: !todo.Completed,
		})
		if err != nil {
			return nil, 500, err
		}
		c.Set("todo", todo)
		return c.Render(`
			<TodoCount filter="all" page="1" />
			<Todo />
		`), 200, nil
	}
	return nil, 404, eris.Errorf("Intent not specified: %s", params.Intent)
}
