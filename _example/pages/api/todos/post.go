package todos

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pyros2097/gromer/_example/db"
)

type PostParams struct {
	Text string `json:"text"`
}

func POST(ctx context.Context, b PostParams) (*db.Todo, int, error) {
	todo, err := db.Query.CreateTodo(ctx, db.CreateTodoParams{
		ID:        uuid.New().String(),
		Text:      b.Text,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, 500, err
	}
	return todo, 200, nil
}
