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

func POST(ctx Context, params PostParams) (*Node, int, error) {
	if params.Intent == "clear_completed" {
		allTodos, err := todos.GetAllTodo(ctx, todos.GetAllTodoParams{
			Filter: "all",
			Limit:  1000,
		})
		if err != nil {
			return nil, 500, err
		}
		for _, t := range allTodos {
			if t.Completed {
				_, err := todos.DeleteTodo(ctx, t.ID)
				if err != nil {
					return nil, 500, err
				}
			}
		}
		return ctx.Render(`
			<TodoList id="todo-list" filter="all" page="1"></TodoList>
			<TodoCount filter="all" page="1"></TodoCount>
		`), 200, nil
	} else if params.Intent == "create" {
		todo, err := todos.CreateTodo(ctx, params.Text)
		if err != nil {
			return nil, 500, err
		}
		ctx.Set("todo", todo)
		return ctx.Render(`
			<Todo todo=todo></Todo>
			<TodoCount filter="all" page="1"></TodoCount>
		`), 200, nil
	} else if params.Intent == "delete" {
		_, err := todos.DeleteTodo(ctx, params.ID)
		if err != nil {
			return nil, 500, err
		}
		return nil, 200, nil
	} else if params.Intent == "complete" {
		todo, err := todos.GetTodo(ctx, params.ID)
		if err != nil {
			return nil, 500, err
		}
		_, err = todos.UpdateTodo(ctx, params.ID, todos.UpdateTodoParams{
			Text:      todo.Text,
			Completed: !todo.Completed,
		})
		if err != nil {
			return nil, 500, err
		}
		ctx.Set("todo", todo)
		return ctx.Render(`
			{{#Todo todo=todo}}{{/Todo}}
		`), 200, nil
	}
	return nil, 404, fmt.Errorf("Intent not specified: %s", params.Intent)
}
