package gsx

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"runtime"
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
	htmlTags = []string{"a", "abbr", "acronym", "address", "applet", "area", "article", "aside", "audio", "b", "base", "basefont", "bb", "bdo", "big", "blockquote", "body", "br /", "button", "canvas", "caption", "center", "cite", "code", "col", "colgroup", "command", "datagrid", "datalist", "dd", "del", "details", "dfn", "dialog", "dir", "div", "dl", "dt", "em", "embed", "eventsource", "fieldset", "figcaption", "figure", "font", "footer", "form", "frame", "frameset", "h1 to <h6>", "head", "header", "hgroup", "hr /", "html", "i", "iframe", "img", "input", "ins", "isindex", "kbd", "keygen", "label", "legend", "li", "link", "map", "mark", "menu", "meta", "meter", "nav", "noframes", "noscript", "object", "ol", "optgroup", "option", "output", "p", "param", "pre", "progress", "q", "rp", "rt", "ruby", "s", "samp", "script", "section", "select", "small", "source", "span", "strike", "strong", "style", "sub", "sup", "table", "tbody", "td", "textarea", "tfoot", "th", "thead", "time", "title", "tr", "track", "tt", "u", "ul", "var", "video", "wbr"}
	compMap  = map[string]ComponentFunc{}
	funcMap  = map[string]interface{}{}
	styles   = ""
	refRegex = regexp.MustCompile(`{(.*?)}`)
)

type (
	ComponentFunc struct {
		Func interface{}
		Args []string
	}
	Html map[string]interface{}
	Node struct {
		*html.Node
	}
)

func (h Html) Render(tpl string) Node {
	newTpl := stripWhitespace(tpl)
	doc, err := html.ParseFragment(bytes.NewBuffer([]byte(newTpl)), contextNode)
	if err != nil {
		panic(err)
	}
	populate(h, doc[0])
	return Node{doc[0]}
}

func (n Node) Write(w io.Writer) error {
	return html.Render(w, n.Node)
}

func (n Node) String() string {
	b := bytes.NewBuffer(nil)
	html.Render(b, n.Node)
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

func RegisterComponent(f interface{}, args ...string) {
	name := strings.ToLower(getFunctionName(f))
	assertName("component", name)
	compMap[name] = ComponentFunc{
		Func: f,
		Args: args,
	}
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
	assertName("function", name)
	funcMap[name] = f
}

func Css(v string) string {
	styles += v
	return v
}

func GetStyles() string {
	return styles
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
	case string:
		return iv
	}
	return nil
}

func getRefValue(ctx map[string]interface{}, ref string) interface{} {
	if f, ok := funcMap[ref]; ok {
		return f.(func() string)()
	} else {
		parts := strings.Split(strings.ReplaceAll(ref, "!", ""), ".")
		if len(parts) == 2 {
			if v, ok := ctx[parts[0]]; ok {
				i := reflect.ValueOf(v).Elem().FieldByName(parts[1]).Interface()
				return convert(ref, i)
			}
		}
		return convert(ref, ctx[ref])
	}
}

func removeBrackets(v string) string {
	return strings.ReplaceAll(strings.ReplaceAll(v, "{", ""), "}", "")
}

func substituteString(ctx map[string]interface{}, v string) string {
	found := refRegex.FindString(v)
	if found != "" {
		varName := removeBrackets(found)
		varValue := fmt.Sprintf("%v", getRefValue(ctx, varName))
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

func populate(ctx Html, n *html.Node) {
	if n.Type == html.TextNode {
		if n.Data != "" && strings.Contains(n.Data, "{") && n.Data != "{children}" {
			n.Data = substituteString(ctx, n.Data)
		}
	} else if n.Type == html.ElementNode {
		for i, at := range n.Attr {
			if at.Key == "x-for" {
				xfor := getAttribute("x-for", n.Attr)
				arr := strings.Split(xfor, " in ")
				ctxItemKey := arr[0]
				ctxKey := arr[1]
				data := ctx[ctxKey]
				switch reflect.TypeOf(data).Kind() {
				case reflect.Slice:
					v := reflect.ValueOf(data)
					firstChild := cloneNode(n.FirstChild)
					n.RemoveChild(n.FirstChild)
					for i := 0; i < v.Len(); i++ {
						compCtx := map[string]interface{}{
							ctxItemKey: v.Index(i).Interface(),
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
				if at.Key == "class" {
					classes := ""
					kvstrings := strings.Split(strings.TrimSpace(removeBrackets(at.Val)), ",")
					for _, kv := range kvstrings {
						kvarray := strings.Split(kv, ":")
						k := strings.TrimSpace(kvarray[0])
						v := strings.TrimSpace(kvarray[1])
						varValue := getRefValue(ctx, v)
						if varValue.(bool) {
							classes += k
						}
					}
					n.Attr[i] = html.Attribute{
						Namespace: at.Namespace,
						Key:       at.Key,
						Val:       classes,
					}
				} else {
					n.Attr[i] = html.Attribute{
						Namespace: at.Namespace,
						Key:       at.Key,
						Val:       substituteString(ctx, at.Val),
					}
				}
			}
		}
		if comp, ok := compMap[n.Data]; ok {
			newNode := populateComponent(ctx, comp, n, true)
			n.AppendChild(newNode)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			populate(ctx, c)
		}
	}
}

func renderComponent(ctx Html, comp ComponentFunc, n *html.Node) Node {
	h := Html(ctx)
	args := []reflect.Value{reflect.ValueOf(h)}
	for _, arg := range comp.Args {
		if v, ok := ctx[arg]; ok {
			args = append(args, reflect.ValueOf(v))
		} else {
			v := getAttribute(arg, n.Attr)
			args = append(args, reflect.ValueOf(v))
		}
	}
	result := reflect.ValueOf(comp.Func).Call(args)
	compNode := result[0].Interface().(Node)
	return compNode
}

func populateComponent(ctx Html, comp ComponentFunc, n *html.Node, remove bool) *html.Node {
	compNode := renderComponent(ctx, comp, n)
	if n.FirstChild != nil {
		newChild := cloneNode(n.FirstChild)
		newChild.Parent = nil
		if n.FirstChild != nil && remove {
			n.RemoveChild(n.FirstChild)
		}
		if !remove {
			populate(ctx, newChild)
		}
		populateChildren(compNode.FirstChild, newChild)
	}
	return compNode.Node
}
