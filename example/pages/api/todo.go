package pages

import (
	"context"
)

type Todo struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func GetParam(c context.Context, k string) string {
	return ""
}

//wapp:api method=GET path=/api/todos
func GetAllTodos(c context.Context) (interface{}, error) {
	return []Todo{}, nil
}

// GetTodoById(c context.Context, id string)

//wapp:api method=GET path=/api/todos/:id
func GetTodoById(c context.Context) (interface{}, error) {
	id := GetParam(c, ":id")
	return Todo{ID: id}, nil
}

// UpdateTodo(c context.Context, id string, body *Todo)

//wapp:api method=POST path=/api/todos
func CreateTodo(c context.Context) (interface{}, error) {
	var todo Todo
	// GetBody(c, &todo)
	return todo, nil
}

//wapp:api method=PUT path=/api/todos/:id
func UpdateTodo(c context.Context) (interface{}, error) {
	id := GetParam(c, ":id")
	return Todo{ID: id}, nil
}
