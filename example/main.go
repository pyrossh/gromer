package main

import (
	app "github.com/pyros2097/wapp"
)

var routes = map[string]app.RenderFunc{
	"/about":     About,
	"/clock":     Clock,
	"/container": Container,
	"/":          Index,
}

func main() {
	app.Run(false, routes)
}
