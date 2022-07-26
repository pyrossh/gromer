package containers

import "github.com/pyros2097/gromer/gsx"

func init() {
	gsx.RegisterComponent(TodoCount, nil, "filter")
	gsx.RegisterComponent(TodoList, nil, "page", "filter")
}
