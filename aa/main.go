package main

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

type M map[string]interface{}

type ReqContext struct {
	JS  *bytes.Buffer
	CSS *bytes.Buffer
}

type Component func(ReqContext) string

var components = map[string]Component{}

func RegisterComponent(k string, v Component) {
	components[k] = v
}

func Html2(ctx ReqContext, input string, data map[string]interface{}) string {
	return input
}

func ParseHTML(ctx ReqContext, input string, data map[string]interface{}) *html.Node {
	doc, err := html.Parse(bytes.NewBufferString(input))
	if err != nil {
		panic(err)
	}
	return doc
}

func init() {
	RegisterComponent("Layout", func(c ReqContext) string {
		return Html2(c, `
		<html>
			<body>
				<div>
				<slot></slot>
				</div>
			</body>
		</html>
		`, M{})
	})
}

func properTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func main() {
	// vctx := velvet.NewContext()
	// for k, v := range data {
	// 	vctx.Set(k, v)
	// }
	// s, err := velvet.Render(textOutput, vctx)
	// if err != nil {
	// 	panic(err)
	// }
	ctx := ReqContext{JS: bytes.NewBuffer(nil), CSS: bytes.NewBuffer(nil)}
	txt := `
		<Layout>
			<p>
				Hello world
			</p>
		</Layout>
	`
	docs := ParseHTML(ctx, txt, M{})
	textOutput := ""
	var f func(txt string, n *html.Node)
	f = func(txt string, n *html.Node) {
		removeExtraStuff := func(tag string) bool {
			return !strings.Contains(txt, "<"+tag+">") && n.Type == html.ElementNode && n.Data == tag
		}
		constainsHtml := removeExtraStuff("html") || removeExtraStuff("body") || removeExtraStuff("head")
		if !constainsHtml && n.Type == html.ElementNode {
			textOutput += "<" + n.Data + ">"
		}
		if n.Type == html.TextNode {
			textOutput += n.Data
		}
		if n.Type == html.ElementNode && n.Data == "slot" {
		}
		if c, ok := components[properTitle(n.Data)]; ok {
			ctext := c(ctx)
			newNodes := ParseHTML(ctx, ctext, M{})
			f(ctext, newNodes)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(txt, c)
		}
		if !constainsHtml && n.Type == html.ElementNode {
			textOutput += "</" + n.Data + ">"
		}
	}
	f(txt, docs)
	println("textOutput", textOutput)
}

// func Index2(c *context.ReqContext) (interface{}, int, error) {
// 	data := M{
// 		"userID":  c.UserID,
// 		"message": "I ❤️ Alpine",
// 	}
// return Html(`
// <page x-data="pageData">
// 	<div class="flex flex-col items-center justify-center">
// 		<header></header>
// 		<h1>Hello <template x-text="userID"></template></h1>
// 		<h2>Hello this is a h1</h1>
// 		<h2>Hello this is a h2</h1>
// 		<h3 x-text="message"></h3>
// 		<counter start={4}></counter>
// 	</div>
// </page>
// 	`, data), 200, nil
// }
