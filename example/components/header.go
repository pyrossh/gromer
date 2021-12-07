package components

import (
	. "github.com/pyros2097/gromer"
)

func Header() *Element {
	link := "border-b-2 border-white hover:border-red-700 mr-4"
	return Row(Css("w-full mb-20 font-bold text-xl text-gray-700 p-4"),
		Div(Css("text-blue-700"),
			A(Href("https://pyros.sh"), Text("pyros.sh")),
		),
		Div(Css("flex flex-row flex-1 justify-end items-end p-2"),
			Div(Css("border-b-2 border-white text-lg text-blue-700 mr-4"), Text("Examples: ")),
			Div(Css(link), A(Href("/"), Text("Home"))),
			Div(Css(link), A(Href("/clock"), Text("Clock"))),
			Div(Css(link), A(Href("/about"), Text("About"))),
			Div(Css(link), A(Href("/container"), Text("Container"))),
			Div(Css(link), A(Href("/panic"), Text("Panic"))),
		),
	)
}