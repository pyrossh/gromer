package main

import (
	. "github.com/pyros2097/wapp"
)

func Route(c *RenderContext) UI {
	return Div(Text("About Me"))
}

func main() {
	Run(Route)
}
