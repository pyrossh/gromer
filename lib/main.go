package main

import (
	"bytes"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type HtmlTag struct {
	Tag      string
	Text     string
	Attr     []html.Attribute
	Children []*HtmlTag
}

var contextNode = &html.Node{
	Type:     html.ElementNode,
	Data:     "div",
	DataAtom: atom.Lookup([]byte("div")),
}

func Html(tpl string) *html.Node {
	newTpl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tpl, "\n", ""), "\r", ""), "\t", "")
	doc, err := html.ParseFragment(bytes.NewBuffer([]byte(newTpl)), contextNode)
	if err != nil {
		panic(err)
	}
	return doc[0]
}

func Todo() *html.Node {
	return Html(`
		<li id="{todo.ID}" class="{ completed: todo.Completed }">
			<div class="view">
				<span>{todo.Text}</span>
			</div>
			{children}
			<div class="count">
				<span>{todo.Count}</span>
			</div>
		</li>
	`)
}

func walkChildren(n, replaceNode1 *html.Node) {
	if n.Data == "{children}" { // first
		replaceNode1.NextSibling = &html.Node{}
		*replaceNode1.NextSibling = *n.NextSibling
		n.Parent.FirstChild = replaceNode1
		return
	}
	if n.NextSibling != nil {
		if n.NextSibling.Data == "{children}" {
			if n.NextSibling.NextSibling != nil { // middle
				replaceNode1.NextSibling = &html.Node{}
				*replaceNode1.NextSibling = *n.NextSibling.NextSibling
				n.NextSibling = replaceNode1
			}
			if n.NextSibling.PrevSibling != nil { // last
				replaceNode1.PrevSibling = &html.Node{}
				*replaceNode1.PrevSibling = *n.NextSibling.PrevSibling
				n.NextSibling = replaceNode1
			}
		} else {
			walkChildren(n.NextSibling, replaceNode1)
		}
	}
	if n.FirstChild != nil {
		walkChildren(n.FirstChild, replaceNode1)
	}
}

func main() {
	// ctx := map[string]interface{}{}
	doc := Html(`
		<div>
			<div>
				123
				<Todo id={todo.ID} class="{ completed: todo.Completed }">
					<div class="container">
						<h2>Title</h2>
						<h3>Sub title</h3>
					</div>
				</Todo>
			</div>
			<div>
				Test
				<button>click</button>
			</div>
		</div>
	`)
	top := &HtmlTag{}
	var f func(parent *HtmlTag, n *html.Node)
	f = func(parent *HtmlTag, n *html.Node) {
		if n.Type == html.TextNode {
			data := strings.TrimSpace(n.Data)
			if data != "" {
				// data = "{}"
			}
		} else if n.Type == html.ElementNode {
			// repr.Println(n.Data)
			if n.Data == "todo" {
				componentNode := Todo()
				html.Render(os.Stdout, componentNode)
				println("")
				html.Render(os.Stdout, n)
				println("")
				if n.FirstChild != nil {
					newChild := &html.Node{}
					*newChild = *n.FirstChild
					newChild.Parent = nil
					n.RemoveChild(n.FirstChild)
					walkChildren(componentNode.FirstChild, newChild)
					n.AppendChild(componentNode)
				}
			}
			// newParent := &HtmlTag{
			// 	Tag:  n.Data,
			// 	Attr: n.Attr,
			// }
			// parent.Children = append(parent.Children, newParent)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(nil, c)
			}
		}
	}
	f(top, doc)
	// repr.Println(top)
	html.Render(os.Stdout, doc)
}
