package pages

import (
	"context"

	. "github.com/pyros2097/gromer"
	. "github.com/pyros2097/gromer/example/components"
	. "github.com/pyros2097/gromer/example/context"
)

func GET(c context.Context) (HtmlPage, int, error) {
	ctx := WithState(c)
	userID := GetUserID(c)
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
				H1(Text("Hello	 "+userID)),
				H1(Text("Hello this is a h1")),
				H2(Text("Hello this is a h2")),
				H3(XData("{ message: 'I ❤️ Alpine' }"), XText("message"), Text("")),
				Div(Css("mt-10"),
					Counter(ctx, 4),
				),
			),
		),
	), 200, nil
}
