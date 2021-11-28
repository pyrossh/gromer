package about

import (
	"context"

	. "github.com/pyros2097/gromer"
	. "github.com/pyros2097/gromer/example/components"
)

func GET(c context.Context) (HtmlPage, int, error) {
	return Html(
		Head(
			Title("Example"),
			Meta("description", "Example"),
			Meta("author", "pyros.sh"),
			Meta("keywords", "pyros.sh, gromer"),
			Meta("viewport", "width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover"),
			Link("icon", "/assets/icon.png"),
			Script(Src("/assets/alpine.js"), Defer()),
		),
		Body(
			Col(
				Header(),
				Row(Css("text-5xl text-gray-700"),
					Text("About Me"),
				),
			),
		),
	), 200, nil
}
