package pages

import (
	"context"
	"fmt"

	_ "github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/handlebars"
)

type PostParams struct {
	Intent string `json:"intent"`
	ID     string `json:"id"`
	Text   string `json:"text"`
}

func POST(ctx context.Context, params PostParams) (HtmlContent, int, error) {
	if params.Intent == "clear_completed" {
		allTodos, err := todos.GetAllTodo(ctx, todos.GetAllTodoParams{
			Filter: "all",
			Limit:  1000,
		})
		if err != nil {
			return HtmlErr(500, err)
		}
		for _, t := range allTodos {
			_, err := todos.DeleteTodo(ctx, t.ID)
			if err != nil {
				return HtmlErr(500, err)
			}
		}
		return Html(`
			{{#TodoList id="todo-list"}}{{/TodoList}}
		`).Render()
	} else if params.Intent == "create" {
		todo, err := todos.CreateTodo(ctx, params.Text)
		if err != nil {
			return HtmlErr(500, err)
		}
		return Html(`
			{{#Todo todo=todo}}{{/Todo}}
			{{#TodoCount filter="all" page=1}}{{/TodoCount}}
		`).
			Prop("todo", todo).
			Render()
	} else if params.Intent == "delete" {
		_, err := todos.DeleteTodo(ctx, params.ID)
		if err != nil {
			return HtmlErr(500, err)
		}
		return HtmlEmpty()
	} else if params.Intent == "complete" {
		todo, err := todos.GetTodo(ctx, params.ID)
		if err != nil {
			return HtmlErr(500, err)
		}
		_, err = todos.UpdateTodo(ctx, params.ID, todos.UpdateTodoParams{
			Text:      todo.Text,
			Completed: !todo.Completed,
		})
		if err != nil {
			return HtmlErr(500, err)
		}
		return Html(`
			{{#Todo todo=todo}}{{/Todo}}
		`).
			Prop("todo", todo).
			Render()
	}
	return HtmlErr(404, fmt.Errorf("Intent not specified: %s", params.Intent))
}
