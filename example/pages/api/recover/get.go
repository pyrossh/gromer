package recover

import (
	"context"
	"fmt"
)

type Params struct {
	Limit int `json:"limit"`
}

func GET(ctx context.Context, params Params) (*Params, int, error) {
	arr := []string{}
	v := arr[55]
	fmt.Printf("%s", v)
	return &params, 200, nil
}
