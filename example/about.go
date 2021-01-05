package main

import (
	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
)

func About(c *RenderContext) UI {
	return Col(
		Header(c),
		Row(Css("text-5xl text-gray-700"),
			HelmetTitle("wapp-example"),
			HelmetDescription("wapp is a framework"),
			HelmetAuthor("pyros2097"),
			HelmetKeywords("wapp,wapp-example,golang,framework,frontend,ui,wasm,isomorphic"),
			Text("About Me"),
		),
	)
}
