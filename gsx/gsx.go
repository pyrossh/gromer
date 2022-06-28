package gsx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	contextNode = &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Lookup([]byte("div")),
	}
	htmlTags   = []string{"a", "abbr", "acronym", "address", "applet", "area", "article", "aside", "audio", "b", "base", "basefont", "bb", "bdo", "big", "blockquote", "body", "br /", "button", "canvas", "caption", "center", "cite", "code", "col", "colgroup", "command", "datagrid", "datalist", "dd", "del", "details", "dfn", "dialog", "dir", "div", "dl", "dt", "em", "embed", "eventsource", "fieldset", "figcaption", "figure", "font", "footer", "form", "frame", "frameset", "h1 to <h6>", "head", "header", "hgroup", "hr /", "html", "i", "iframe", "img", "input", "ins", "isindex", "kbd", "keygen", "label", "legend", "li", "link", "map", "mark", "menu", "meta", "meter", "nav", "noframes", "noscript", "object", "ol", "optgroup", "option", "output", "p", "param", "pre", "progress", "q", "rp", "rt", "ruby", "s", "samp", "script", "section", "select", "small", "source", "span", "strike", "strong", "style", "sub", "sup", "table", "tbody", "td", "textarea", "tfoot", "th", "thead", "time", "title", "tr", "track", "tt", "u", "ul", "var", "video", "wbr"}
	compMap    = map[string]ComponentFunc{}
	funcMap    = map[string]interface{}{}
	classesMap = map[string]M{}
	refRegex   = regexp.MustCompile(`{(.*?)}`)
)

type (
	M             map[string]interface{}
	MS            map[string]string
	Arr           []interface{}
	ComponentFunc struct {
		Func    interface{}
		Args    []string
		Classes M
	}
	link struct {
		Rel  string
		Href string
		Type string
		As   string
	}
	Context struct {
		context.Context
		hxRequest bool
		data      M
		meta      M
		links     map[string]link
		scripts   map[string]bool
	}
	Node struct {
		html.Node
	}
)

func NewContext(c context.Context, hxRequest bool) *Context {
	return &Context{
		Context:   c,
		hxRequest: hxRequest,
		data:      M{},
		meta:      M{},
		links:     map[string]link{},
		scripts:   map[string]bool{},
	}
}

func (c *Context) Get(k string) interface{} {
	return c.data[k]
}

func (c *Context) Set(k string, v interface{}) {
	c.data[k] = v
}

func (c *Context) Meta(meta M) {
	c.meta = meta
}

func (c *Context) AddMeta(k, v string) {
	c.meta[k] = v
}

func (c *Context) Link(rel, href, t, as string) {
	c.links[href] = link{rel, href, t, as}
}

func (c *Context) Script(src string, sdefer bool) {
	c.scripts[src] = sdefer
}

func (c *Context) Data(data M) {
	c.data = data
}

func (c *Context) Render(tpl string) *Node {
	newTpl := stripWhitespace(tpl)
	doc, err := html.ParseFragment(bytes.NewBuffer([]byte(newTpl)), contextNode)
	if err != nil {
		panic(err)
	}
	populate(c, doc[0])
	return &Node{*doc[0]}
}

func (n *Node) Write(c *Context, w io.Writer) {
	if !c.hxRequest {
		w.Write([]byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">`))
		w.Write([]byte(`<meta http-equiv="Content-Type" content="text/html;charset=utf-8"><meta content="utf-8" http-equiv="encoding">`))
		w.Write([]byte(`<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover">`))
		for k, v := range c.meta {
			w.Write([]byte(fmt.Sprintf(`<meta name="%s" content="%s">`, k, v)))
		}
		for k, v := range c.meta {
			if k == "title" {
				w.Write([]byte(fmt.Sprintf(`<title>%s</title>`, v)))
			}
		}
		for _, v := range c.links {
			if v.Type != "" || v.As != "" {
				w.Write([]byte(fmt.Sprintf(`<link rel="%s" href="%s" type="%s" as="%s">`, v.Rel, v.Href, v.Type, v.As)))
			} else {
				w.Write([]byte(fmt.Sprintf(`<link rel="%s" href="%s">`, v.Rel, v.Href)))
			}
		}
		for src, sdefer := range c.scripts {
			if sdefer {
				w.Write([]byte(fmt.Sprintf(`<script src="%s" defer="true"></script>`, src)))
			} else {
				w.Write([]byte(fmt.Sprintf(`<script src="%s"></script>`, src)))
			}
		}
		w.Write([]byte(`</head><body>`))
	}
	html.Render(w, &n.Node)
	if !c.hxRequest {
		w.Write([]byte(`</body></html>`))
	}
}

func (n *Node) String() string {
	b := bytes.NewBuffer(nil)
	html.Render(b, &n.Node)
	return b.String()
}

func stripWhitespace(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), "\r", ""), "\t", "")
}

func assertName(t, name string) {
	for _, v := range htmlTags {
		if name == v {
			panic(fmt.Sprintf("%s '%s' name cannot be the same as a html tag", t, name))
		}
	}
}

func SetClasses(k string, m M) {
	classesMap[k] = m
}

func GetPageStyles(k string) string {
	return normalizeCss + "\n" + computeCss(classesMap[k], k)
}

func GetComponentStyles() string {
	css := ""
	for k, v := range compMap {
		if v.Classes != nil {
			css += computeCss(v.Classes, k)
		}
	}
	return css
}

func RegisterComponent(f interface{}, classes M, args ...string) {
	name := strings.ToLower(getFunctionName(f))
	assertName("component", name)
	compMap[name] = ComponentFunc{
		Func:    f,
		Args:    args,
		Classes: classes,
	}
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
	assertName("function", name)
	funcMap[name] = f
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func getAttribute(k string, kvs []html.Attribute) string {
	for _, v := range kvs {
		if v.Key == k {
			return v.Val
		}
	}
	return ""
}

func convert(ref string, i interface{}) interface{} {
	switch iv := i.(type) {
	case bool:
		if strings.Contains(ref, "!") {
			return !iv
		} else {
			return iv
		}
	case int:
		return iv
	case string:
		return iv
	default:
		return iv
	}
}

func getRefValue(c *Context, ref string) interface{} {
	if ref == "true" {
		return true
	} else if ref == "false" {
		return false
	} else if f, ok := funcMap[ref]; ok {
		return f.(func() string)()
	} else {
		parts := strings.Split(strings.ReplaceAll(ref, "!", ""), ".")
		if len(parts) == 2 {
			if v, ok := c.data[parts[0]]; ok {
				a := reflect.ValueOf(v)
				if a.Kind() == reflect.Ptr {
					i := a.Elem().FieldByName(parts[1]).Interface()
					return convert(ref, i)
				} else {
					i := a.FieldByName(parts[1]).Interface()
					return convert(ref, i)
				}
			}
		}
		return convert(ref, c.data[ref])
	}
}

func removeBrackets(v string) string {
	return strings.ReplaceAll(strings.ReplaceAll(v, "{", ""), "}", "")
}

func substituteString(c *Context, v string) string {
	found := refRegex.FindString(v)
	if found != "" {
		varName := removeBrackets(found)
		varValue := fmt.Sprintf("%v", getRefValue(c, varName))
		return strings.ReplaceAll(v, found, varValue)
	}
	return v
}

func populateChildren(n, replaceNode1 *html.Node) {
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
			populateChildren(n.NextSibling, replaceNode1)
		}
	}
	if n.FirstChild != nil {
		populateChildren(n.FirstChild, replaceNode1)
	}
}

func cloneNode(n *html.Node) *html.Node {
	attrs := []html.Attribute{}
	for _, v := range n.Attr {
		attrs = append(attrs, html.Attribute{
			Key: v.Key,
			Val: v.Val,
		})
	}
	newNode := &html.Node{
		Type:     n.Type,
		Data:     n.Data,
		DataAtom: n.DataAtom,
		Attr:     attrs,
	}
	if n.FirstChild != nil {
		newNode.FirstChild = cloneNode(n.FirstChild)
	}
	if n.NextSibling != nil {
		newNode.NextSibling = cloneNode(n.NextSibling)
	}
	return newNode
}

func populate(c *Context, n *html.Node) {
	if n.Type == html.TextNode {
		if n.Data != "" && strings.Contains(n.Data, "{") && n.Data != "{children}" {
			n.Data = substituteString(c, n.Data)
		}
	} else if n.Type == html.ElementNode {
		for i, at := range n.Attr {
			if at.Key == "x-for" {
				xfor := getAttribute("x-for", n.Attr)
				arr := strings.Split(xfor, " in ")
				ctxItemKey := arr[0]
				ctxKey := arr[1]
				data := c.data[ctxKey]
				switch reflect.TypeOf(data).Kind() {
				case reflect.Slice:
					v := reflect.ValueOf(data)
					if v.Len() == 0 {
						if n.FirstChild != nil {
							n.RemoveChild(n.FirstChild)
						}
						continue
					}
					if n.FirstChild == nil {
						continue
					}
					firstChild := cloneNode(n.FirstChild)
					n.RemoveChild(n.FirstChild)
					for i := 0; i < v.Len(); i++ {
						compCtx := &Context{
							Context: c.Context,
							data: map[string]interface{}{
								ctxItemKey: v.Index(i).Interface(),
							},
						}
						itemChild := cloneNode(firstChild)
						itemChild.Parent = nil
						if comp, ok := compMap[itemChild.Data]; ok {
							newNode := populateComponent(compCtx, comp, itemChild, false)
							n.AppendChild(newNode)
						} else {
							n.AppendChild(itemChild)
							populate(compCtx, itemChild)
						}
					}
				}
			} else if at.Val != "" && strings.Contains(at.Val, "{") {
				if at.Key == "class" || at.Key == "src" {
					classes := []string{}
					kvstrings := strings.Split(strings.TrimSpace(removeBrackets(at.Val)), ",")
					for _, kv := range kvstrings {
						kvarray := strings.Split(kv, ":")
						k := strings.TrimSpace(kvarray[0])
						v := strings.TrimSpace(kvarray[1])
						varValue := getRefValue(c, v)
						if varValue.(bool) {
							classes = append(classes, k)
						}
					}
					n.Attr[i] = html.Attribute{
						Namespace: at.Namespace,
						Key:       at.Key,
						Val:       strings.Join(classes, " "),
					}
				} else {
					n.Attr[i] = html.Attribute{
						Namespace: at.Namespace,
						Key:       at.Key,
						Val:       substituteString(c, at.Val),
					}
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			populate(c, child)
		}
		if comp, ok := compMap[n.Data]; ok {
			newNode := populateComponent(c, comp, n, true)
			n.AppendChild(newNode)
		}
	}
}

func renderComponent(c *Context, comp ComponentFunc, n *html.Node) *Node {
	args := []reflect.Value{reflect.ValueOf(c)}
	funcType := reflect.TypeOf(comp.Func)
	for i, arg := range comp.Args {
		if v, ok := c.data[arg]; ok {
			args = append(args, reflect.ValueOf(v))
		} else {
			v := getAttribute(arg, n.Attr)
			t := funcType.In(i + 1)
			switch t.Kind() {
			case reflect.Int:
				value, _ := strconv.Atoi(v)
				args = append(args, reflect.ValueOf(value))
			case reflect.Bool:
				value, _ := strconv.ParseBool(v)
				args = append(args, reflect.ValueOf(value))
			default:
				args = append(args, reflect.ValueOf(v))
			}
		}
	}
	result := reflect.ValueOf(comp.Func).Call(args)
	compNode := result[0].Interface().(*Node)
	return compNode
}

func populateComponent(c *Context, comp ComponentFunc, n *html.Node, remove bool) *html.Node {
	compNode := renderComponent(c, comp, n)
	if n.FirstChild != nil {
		newChild := cloneNode(n.FirstChild)
		newChild.Parent = nil
		if n.FirstChild != nil && remove {
			n.RemoveChild(n.FirstChild)
		}
		if !remove {
			populate(c, newChild)
		}
		if compNode.FirstChild != nil {
			populateChildren(compNode.FirstChild, newChild)
		}
	}
	return &compNode.Node
}
