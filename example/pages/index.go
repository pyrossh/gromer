package pages

import (
	"net/http"
	"strconv"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
)

func Index(w http.ResponseWriter, r *http.Request) *Element {
	return Page(
		Col(
			Header(),
			H1(Text("Hello this is a h1")),
			H2(Text("Hello this is a h2")),
			H2(XData("{ message: 'I ❤️ Alpine' }"), XText("message"), Text("")),
			Col(Css("text-3xl text-gray-700"),
				Row(
					Row(Css("underline"),
						Text("Counter"),
					),
				),
				Row(
					Button(Css("btn m-20"),
						Text("-"),
					),
					Row(Css("m-20"),
						Text(strconv.Itoa(1)),
					),
					Button(Css("btn m-20"),
						Text("+"),
					),
				),
			),
		),
	)
}
