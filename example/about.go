package main

import (
	. "github.com/pyros2097/wapp"
)

func About(c *RenderContext) UI {
	return Div(
		HelmetTitle("wapp-example"),
		HelmetDescription("wapp is a framework"),
		HelmetAuthor("pyros2097"),
		HelmetKeywords("wapp,wapp-example,golang,framework,frontend,ui,wasm,isomorphic"),
		Text("About Me"),
	)
}
