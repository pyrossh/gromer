package routes

import (
	. "github.com/pyros2097/gromer/gsx"
)

func AboutPage(c *Context) ([]*Tag, int, error) {
	c.Meta(M{
		"title":       "About Gromer",
		"description": "About Gromer",
	})
	c.Styles(M{
		"container": "flex flex-col justify-center items-center",
	})
	return c.Render(`
		<div class="About">
			A new link is here
			<h1>About Me</h1>
		</div>
	`), 200, nil
}
