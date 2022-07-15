package not_found_404

import (
	. "github.com/pyros2097/gromer/gsx"
)

// var (
// 	Meta = M{
// 		"title": "Oops something wen't wrong",
// 	}
// 	Styles = M{
// 		"notfound": "flex flex-1 items-center justify-center w-screen h-screen",
// 		"back":     "mt-6",
// 		"link":     "underlined",
// 	}
// )

// var (
// 	Meta = M{
// 		"title": "Page Not Found",
// 	}
// 	Styles = M{
// 		"notfound": "flex flex-1 items-center justify-center w-screen h-screen",
// 		"back":     "mt-6",
// 		"link":     "underlined",
// 	}
// )

// func GET(c *Context) ([]*Tag, int, error) {
// 	return c.Render(`
// 		<div id="error" class="notfound	" hx-swap-oob="true">
// 			<h1>"Page Not Found"</h1>
// 			<h2 class="mt-6">
// 				<a class="link" href="/">"Go Back"</a>
// 			</h2>
// 		</div>
// 	`), 404, nil
// }

var StatusStyles = M{
	"container": "border-t-2 border-gray-100 text-2xl",
	"row":       "flex flex-row group",
	"button-1":  "ml-4 text-gray-400",
	"label":     "flex-1 min-w-0 flex items-center break-all ml-2 p-2 text-gray-800",
	"striked":   "text-gray-500 line-through",
	"button-2":  "mr-4 text-red-700 opacity-0 hover:opacity-100",
	"unchecked": "text-gray-200",
}

func Status(c *Context, status int, err error) []*Tag {
	return c.Render(`
		<div id="error" class="notfound	" hx-swap-oob="true">
			<h1>"Oops something wen't wrong"</h1>
			<h2 class="mt-6">
				<a class="link" href="/">"Go Back"</a>
			</h2>
		</div>
	`), 404, nil
}
