package components

import "github.com/pyros2097/gromer/gsx"

func init() {
	gsx.RegisterComponent(Todo, TodoStyles, "todo")
	gsx.RegisterComponent(Status, StatusStyles, "status", "error")
}
