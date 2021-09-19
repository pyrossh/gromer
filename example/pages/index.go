package pages

import (
	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
	. "github.com/pyros2097/wapp/example/context"
)

func init() {
	RegisterComponent("page", func(c ReqContext) string {
		return Html2(c, `
		<html>
			<body>
				<slot></slot>
			</body>
		</html>
		`, M{})
	})
}

func Index(c ReqContext) (interface{}, int, error) {
	ss := Html2(c, `
	<page x-data="pageData">
		<div class="flex flex-col items-center justify-center">
			<header></header>
			<h1>Hello {{ userID }}</h1>
			<h2>Hello this is a h1</h1>
			<h2>Hello this is a h2</h1>
			<h3 x-text="message"></h3>
			<counter start={4}></counter>
		</div>
	</page>
	`, M{"userID": c.UserID})
	println("ss", ss)
	return Page(c,
		Col(
			Header(),
			H1(Text("Hello "+c.UserID)),
			H1(Text("Hello this is a h1")),
			H2(Text("Hello this is a h2")),
			H3(XData("{ message: 'I ❤️ Alpine' }"), XText("message"), Text("")),
			Counter(c, 4),
		),
	), 200, nil
}

// func Index2(c *context.ReqContext) (interface{}, int, error) {
// 	data := M{
// 		"userID":  c.UserID,
// 		"message": "I ❤️ Alpine",
// 	}
// return Html(`
// 	<page x-data="pageData">
// 		<div class="flex flex-col items-center justify-center">
// 			<header></header>
// 			<h1>Hello <template x-text="userID"></template></h1>
// 			<h2>Hello this is a h1</h1>
// 			<h2>Hello this is a h2</h1>
// 			<h3 x-text="message"></h3>
// 			<counter start={4}></counter>
// 		</div>
// 	</page>
// 	`, data), 200, nil
// }
