package todos

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pyros2097/gromer/_example/services"
)

type PostParams struct {
	Text string `json:"text"`
}

func POST(ctx context.Context, b PostParams) (*services.Todo, int, error) {
	todo, err := services.CreateTodo(ctx, services.Todo{
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
