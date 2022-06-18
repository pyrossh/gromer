package gsx

import (
	"bytes"
	"io"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/alecthomas/repr"
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
	newTpl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tpl, "\n", ""), "\r", ""), "\t", "")
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

func RegisterComponent(name string, f interface{}, args ...string) {
	compMap[name] = ComponentFunc{
		Func: f,
		Args: args,
	}
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
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

func subsRef(ctx map[string]interface{}, ref string) interface{} {
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

func populate(ctx Html, n *html.Node) {
	if n.Type == html.TextNode {
		// if n.Data != "" {
		// }
	} else if n.Type == html.ElementNode {
		repr.Println("dd", n.Data)
		for i, at := range n.Attr {
			// if len(param.Value.KV) != 0 {
			// 	values := []string{}
			// 	for _, kv := range param.Value.KV {
			// 		if subsRef(ctx, kv.Value) == true {
			// 			values = append(values, kv.Key)
			// 		}
			// 	}
			if at.Val != "" && strings.Contains(at.Val, "{") {
				if at.Key == "class" {
					repr.Println(at)
				} else {
					re := regexp.MustCompile(`{(.*?)}`)
					found := re.FindString(at.Val)
					if found != "" {
						varName := strings.ReplaceAll(strings.ReplaceAll(found, "{", ""), "}", "")
						varValue := subsRef(ctx, varName).(string)
						n.Attr[i] = html.Attribute{
							Namespace: at.Namespace,
							Key:       at.Key,
							Val:       strings.ReplaceAll(at.Val, found, varValue),
						}
					}
				}
			}
		}
		if comp, ok := compMap[n.Data]; ok {
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
			// html.Render(os.Stdout, componentNode)
			// println("")
			// html.Render(os.Stdout, n)
			// println("")
			if n.FirstChild != nil {
				newChild := &html.Node{}
				*newChild = *n.FirstChild
				newChild.Parent = nil
				n.RemoveChild(n.FirstChild)
				populateChildren(compNode.FirstChild, newChild)
				n.AppendChild(compNode.Node)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			populate(ctx, c)
		}
	}
}

// func render(x *Xml, ctx map[string]interface{}) string {
// 	if x.Value != nil {
// 		if x.Value.Ref != "" {
// 			s += space + "  " + subsRef(ctx, x.Value.Ref).(string) + "\n"
// 		} else if x.Value.Str != "" {
// 			s += space + "  " + strings.ReplaceAll(x.Value.Str, `"`, "") + "\n"
// 		}
// 	}
// 	if x.Name == "For" {
// 		ctxKey := getAttribute("key", x.Attributes)
// 		ctxName := getAttribute("itemKey", x.Attributes)
// 		data := ctx[ctxKey]
// 		switch reflect.TypeOf(data).Kind() {
// 		case reflect.Slice:
// 			v := reflect.ValueOf(data)
// 			for i := 0; i < v.Len(); i++ {
// 				ctx["_space"] = space + "  "
// 				ctx[ctxName] = v.Index(i).Interface()
// 				s += render(x.Children[0], ctx) + "\n"
// 			}
// 		}
// 	} else {
// 		if comp, ok := compMap[x.Name]; ok {
// 		} else {
// 			found := false
// 			for _, t := range htmlTags {
// 				if t == x.Name {
// 					found = true
// 				}
// 			}
// 			if !found {
// 				panic(fmt.Errorf("Comp not found %s", x.Name))
// 			}
// 			for _, c := range x.Children {
// 				ctx["_space"] = space + "  "
// 				s += render(c, ctx) + "\n"
// 			}
// 		}
// 	}
// 	s += space + "</" + x.Name + ">"
// 	return s
// }

// <script>
//     document.addEventListener('alpine:init', () => {
//         Alpine.store('todos', {
//         		list: [],
//						count: 0,
//         })
//     })
// </script>

// patch: {
// 	{ "op": "add", "path": "/todos/list", "value": { "id": "123", "text": "123" } },
// }
//
