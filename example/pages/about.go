package pages

import (
	"net/http"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
)

func About(w http.ResponseWriter, r *http.Request) *Element {
	return Page(
		Col(
			Header(),
			Row(Css("text-5xl text-gray-700"),
				Text("About Me"),
			),
		),
	)
}
