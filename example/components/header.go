package components

import (
	. "github.com/pyros2097/wapp"
)

func Header() *Element {
	return Row(Css("w-full mb-20 font-bold text-xl text-gray-700 p-4"),
		Div(Css("text-blue-700"),
			A(Href("https://wapp.pyros2097.dev"), Text("wapp.pyros2097.dev")),
		),
		Div(Css("flex flex-row flex-1 justify-end items-end p-2"),
			Div(Css("border-b-2 border-white text-lg text-blue-700 mr-4"), Text("Examples: ")),
			Div(Css("link mr-4"), A(Href("/"), Text("Home"))),
			Div(Css("link mr-4"), A(Href("/clock"), Text("Clock"))),
			Div(Css("link mr-4"), A(Href("/about"), Text("About"))),
			Div(Css("link mr-4"), A(Href("/container"), Text("Container"))),
			Div(Css("link mr-4"), A(Href("/panic"), Text("Panic"))),
		),
	)
}
