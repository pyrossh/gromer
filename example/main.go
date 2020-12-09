package main

import (
	. "github.com/pyros2097/wapp"
)

var routes = map[string]RenderFunc{
	"/about":     About,
	"/clock":     Clock,
	"/container": Container,
	"/":          Index,
}

func main() {
	Run(AppInfo{
		Title:       "wapp-example",
		Description: "wapp is a framework",
		Author:      "pyros2097",
		Keywords:    "wapp,wapp-example,golang,framework,frontend,ui,wasm,isomorphic",
	}, routes)
}
