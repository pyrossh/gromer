package main

import (
	. "github.com/pyros2097/wapp"
)

func main() {
	info := RouteInfo{
		Title:       "wapp-example",
		Description: "wapp is a framework",
		Author:      "pyros2097",
		Keywords:    "wapp,wapp-example,golang,framework,frontend,ui,wasm,isomorphic",
	}
	Route("/about", About, info)
	Route("/clock", Clock, info)
	Route("/container", Container, info)
	Route("/", Index, info)
	Run()
}
