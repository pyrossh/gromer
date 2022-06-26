package about

import (
	. "github.com/pyros2097/gromer/gsx"
)

func GET(c Context) (*Node, int, error) {
	c.Meta("title", "About Gromer")
	c.Meta("description", "About Gromer")
	return c.Render(`
		<div class="flex flex-col justify-center items-center">
			A new link is here
			P<h1>About Me</h1>
		</div>
	`), 200, nil
}
