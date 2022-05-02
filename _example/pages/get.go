package pages

import (
	"context"

	_ "github.com/pyros2097/gromer/_example/components"
	. "github.com/pyros2097/gromer/handlebars"
)

type GetParams struct {
}

func GET(ctx context.Context, params GetParams) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="gromer example"}}
			{{#Header}}{{/Header}}
			Home Page
		{{/Page}}
		`).
		Render()
}
