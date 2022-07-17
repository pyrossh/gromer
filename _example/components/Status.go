package components

import (
	. "github.com/pyros2097/gromer/gsx"
)

var StatusStyles = M{
	"container":      "flex flex-1 flex-col items-center justify-center w-screen h-screen",
	"title":          "text-xl font-bold",
	"back-container": "mt-6 text-lg",
	"link":           "underline",
}

func Status(c *Context, status int, err error) []*Tag {
	if status == 404 {
		c.AddMeta("title", "Gromer | Page Not Found")
		return c.Render(`
			<div id="error" class="Status" hx-swap-oob="true">
				<h1 class="title">"Page Not Found"</h1>
				<h2 class="back-container">
					<a class="link" href="/">"Go back"</a>
				</h2>
			</div>
		`)
	}
	c.AddMeta("title", "Gromer | Oop's something went wrong")
	return c.Render(`
		<div id="error" class="Status" hx-swap-oob="true">
			<h1 class="title">"Oop's something went wrong"</h1>
			<h2 class="back-container">
				<a class="link" href="/">"Go back"</a>
			</h2>
		</div>
	`)
}
