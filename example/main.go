package main

import (
	. "github.com/pyros2097/wapp"
)

func main() {
	Route("/about", About)
	Route("/clock", Clock)
	Route("/container", Container)
	Route("/", Index)
	Run()
}
