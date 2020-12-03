// GENERATED FILE DO NOT EDIT
package main

import (
	. "github.com/pyros2097/wapp"
	"github.com/pyros2097/wapp/js"
)

func main() {
	
		if js.Window.URL().Path == "/" {
			Run(Index)
		}
	
		if js.Window.URL().Path == "/about" {
			Run(About)
		}
	
		if js.Window.URL().Path == "/clock" {
			Run(Clock)
		}
	
		if js.Window.URL().Path == "/container" {
			Run(Container)
		}
	
}
