package todos

import (
	"context"

	"github.com/pyros2097/gromer/example/db"
)

// type GetParams struct {
// 	Limit  int `json:"limit"`
// 	Offset int `json:"limit"`
// }
// , params GetParams

func GET(ctx context.Context) ([]*db.Todo, int, error) {
	todos, err := db.Query.ListTodos(ctx, db.ListTodosParams{
		Limit:  int32(10),
		Offset: int32(0),
	})
	if err != nil {
		return nil, 500, err
	}
	return todos, 200, nil
}
