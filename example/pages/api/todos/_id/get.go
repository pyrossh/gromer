package todo

import (
	"context"

	. "github.com/pyros2097/wapp/example/models"
)

func GET(c context.Context, id string, params map[string]interface{}) (interface{}, error) {
	return Todo{}, nil
}
