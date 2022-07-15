package todos

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rotisserie/eris"
)

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var globalTodos = []*Todo{}

func CreateTodo(ctx context.Context, text string) (*Todo, error) {
	todo := &Todo{
		ID:        uuid.New().String(),
		Text:      text,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	globalTodos = append(globalTodos, todo)
	return todo, nil
}

type UpdateTodoParams struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func UpdateTodo(ctx context.Context, id string, params UpdateTodoParams) (*Todo, error) {
	updateIndex := -1
	for i, todo := range globalTodos {
		if todo.ID == id {
			updateIndex = i
		}
	}
	if updateIndex != -1 {
		globalTodos[updateIndex].Text = params.Text
		globalTodos[updateIndex].Completed = params.Completed
		globalTodos[updateIndex].UpdatedAt = time.Now()
		return globalTodos[updateIndex], nil
	}
	return nil, eris.New("Todo not found")
}

func DeleteTodo(ctx context.Context, id string) (string, error) {
	deleteIndex := -1
	for i, todo := range globalTodos {
		if todo.ID == id {
			deleteIndex = i
		}
	}
	if deleteIndex != -1 {
		globalTodos = append(globalTodos[:deleteIndex], globalTodos[deleteIndex+1:]...)
		return id, nil
	}
	return "", eris.New("Todo not found")
}

func GetTodo(ctx context.Context, id string) (*Todo, error) {
	for _, todo := range globalTodos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, eris.New("Todo not found")
}

type GetAllTodoParams struct {
	Limit  int    `json:"limit"`
	Filter string `json:"filter"`
}

func GetAllTodo(ctx context.Context, params GetAllTodoParams) ([]*Todo, error) {
	// limit := Default(params.Limit, 10)
	if params.Filter == "completed" {
		newTodos := []*Todo{}
		for _, v := range globalTodos {
			if v.Completed {
				newTodos = append(newTodos, v)
			}
		}
		return newTodos, nil
	}
	if params.Filter == "active" {
		newTodos := []*Todo{}
		for _, v := range globalTodos {
			if !v.Completed {
				newTodos = append(newTodos, v)
			}
		}
		return newTodos, nil
	}
	return globalTodos, nil
}
