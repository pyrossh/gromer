package pages

import (
	"context"

	. "github.com/pyros2097/wapp/example/models"
)

func GET(c context.Context) (interface{}, error) {
	return []Todo{}, nil
}
