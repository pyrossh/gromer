package main

import (
	"github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/assets"
	"github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/containers"
	"github.com/pyros2097/gromer/_example/routes"
	"github.com/pyros2097/gromer/gsx"
)

func main() {
	gsx.RegisterComponent(components.Todo, components.TodoStyles, "todo")
	gsx.RegisterComponent(components.Status, components.StatusStyles, "status", "error")
	gsx.RegisterComponent(containers.TodoCount, nil, "filter")
	gsx.RegisterComponent(containers.TodoList, nil, "page", "filter")
	gromer.Init(components.Status, assets.FS)
	gromer.PageRoute("/", routes.TodosPage, routes.TodosAction)
	gromer.PageRoute("/about", routes.AboutPage, nil)
	gromer.Run("3000")
}
