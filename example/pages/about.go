package pages

import (
	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
	"github.com/pyros2097/wapp/example/context"
)

//wapp:page method=GET path=/about
func About(c context.ReqContext) (interface{}, int, error) {
	return Page(c,
		Col(
			Header(),
			Row(Css("text-5xl text-gray-700"),
				Text("About Me"),
			),
		),
	), 200, nil
}
