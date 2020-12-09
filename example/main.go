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
	Run(routes)
}
