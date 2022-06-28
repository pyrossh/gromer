package about

import (
	. "github.com/pyros2097/gromer/gsx"
)

var (
	Meta = M{
		"title":       "About Gromer",
		"description": "About Gromer",
	}
	Styles = M{}
)

func GET(c *Context) (*Node, int, error) {
	return c.Render(`
		<div class="flex flex-col justify-center items-center">
			A new link is here
			P<h1>About Me</h1>
		</div>
	`), 200, nil
}
