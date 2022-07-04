package routes

import (
	"fmt"

	_ "github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

type PostParams struct {
	Intent string `json:"intent"`
	ID     string `json:"id"`
	Text   string `json:"text"`
}

func POST(c *Context, params PostParams) ([]*Tag, int, error) {
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
	return nil, 404, fmt.Errorf("Intent not specified: %s", params.Intent)
}
