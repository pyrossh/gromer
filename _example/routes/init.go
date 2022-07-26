package routes

import (
	"github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/assets"
)

func init() {
	gromer.Init(components.Status, assets.FS)
	gromer.PageRoute("/", TodosPage, TodosAction)
	gromer.PageRoute("/about", AboutPage, nil)
}
