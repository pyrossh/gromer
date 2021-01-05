package main

import (
	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
)

func main() {
	AppRouter.Error = func(c *RenderContext, err error) UI {
		return Col(Css("text-4xl text-gray-700"),
			Header(c),
			Row(
				Text("Oops something went wrong"),
			),
			Row(Css("mt-6"),
				Text("Please check back again"),
			),
		)
	}
	Route("/panic", Panic)
	Route("/about", About)
	Route("/clock", Clock)
	Route("/container", Container)
	Route("/", Index)
	Run()
}
