package gsx

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type ComponentFunc struct {
	Func interface{}
	Args []string
}

type Html map[string]interface{}

func (h Html) Render(tpl string) string {
	tree := &Module{}
	err := xmlParser.ParseBytes(tpl[0:10], []byte(tpl), tree)
	if err != nil {
		panic(err)
	}
	o := ""
	for _, n := range tree.Nodes {
		v := render(n, h)
		o += v
	}
	return o
}

var htmlTags = []string{"a", "abbr", "acronym", "address", "applet", "area", "article", "aside", "audio", "b", "base", "basefont", "bb", "bdo", "big", "blockquote", "body", "br /", "button", "canvas", "caption", "center", "cite", "code", "col", "colgroup", "command", "datagrid", "datalist", "dd", "del", "details", "dfn", "dialog", "dir", "div", "dl", "dt", "em", "embed", "eventsource", "fieldset", "figcaption", "figure", "font", "footer", "form", "frame", "frameset", "h1 to <h6>", "head", "header", "hgroup", "hr /", "html", "i", "iframe", "img", "input", "ins", "isindex", "kbd", "keygen", "label", "legend", "li", "link", "map", "mark", "menu", "meta", "meter", "nav", "noframes", "noscript", "object", "ol", "optgroup", "option", "output", "p", "param", "pre", "progress", "q", "rp", "rt", "ruby", "s", "samp", "script", "section", "select", "small", "source", "span", "strike", "strong", "style", "sub", "sup", "table", "tbody", "td", "textarea", "tfoot", "th", "thead", "time", "title", "tr", "track", "tt", "u", "ul", "var", "video", "wbr"}
var compMap = map[string]ComponentFunc{}
var funcMap = map[string]interface{}{}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RegisterComponent(f interface{}, args ...string) {
	name := getFunctionName(f)
	compMap[name] = ComponentFunc{
		Func: f,
		Args: args,
	}
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
	funcMap[name] = f
}

var styles = ""

func Css(v string) string {
	styles += v
	return v
}

func GetStyles() string {
	return styles
}

func getAttribute(k string, kvs []*Attribute) string {
	for _, param := range kvs {
		if param.Key == k {
			return strings.ReplaceAll(param.Value.Str, `"`, "")
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

func render(x *Xml, ctx map[string]interface{}) string {
	space, _ := ctx["_space"].(string)
	s := space + "<" + x.Name
	if len(x.Attributes) > 0 {
		s += " "
	}
	for i, param := range x.Attributes {
		if len(param.Value.KV) != 0 {
			values := []string{}
			for _, kv := range param.Value.KV {
				if subsRef(ctx, kv.Value) == true {
					values = append(values, kv.Key)
				}
			}
			s += param.Key + `="` + strings.Join(values, "") + `"`
		} else if param.Value.Ref != "" {
			s += param.Key + `="` + subsRef(ctx, param.Value.Ref).(string) + `"`
		} else {
			s += param.Key + "=" + param.Value.Str
		}
		if i < len(x.Attributes)-1 {
			s += " "
		}
	}
	s += ">\n"
	if x.Value != nil {
		if x.Value.Ref != "" {
			s += space + "  " + subsRef(ctx, x.Value.Ref).(string) + "\n"
		} else if x.Value.Str != "" {
			s += space + "  " + strings.ReplaceAll(x.Value.Str, `"`, "") + "\n"
		}
	}
	if x.Name == "For" {
		ctxKey := getAttribute("key", x.Attributes)
		ctxName := getAttribute("itemKey", x.Attributes)
		data := ctx[ctxKey]
		switch reflect.TypeOf(data).Kind() {
		case reflect.Slice:
			v := reflect.ValueOf(data)
			for i := 0; i < v.Len(); i++ {
				ctx["_space"] = space + "  "
				ctx[ctxName] = v.Index(i).Interface()
				s += render(x.Children[0], ctx) + "\n"
			}
		}
	} else {
		if comp, ok := compMap[x.Name]; ok {
			ctx["_space"] = space + "  "
			h := Html(ctx)
			args := []reflect.Value{reflect.ValueOf(h)}
			for _, k := range comp.Args {
				if v, ok := ctx[k]; ok {
					args = append(args, reflect.ValueOf(v))
				} else {
					v := getAttribute(k, x.Attributes)
					args = append(args, reflect.ValueOf(v))
				}
			}
			if len(x.Children) > 0 {
				h["children"] = render(x.Children[0], h)
			}
			result := reflect.ValueOf(comp.Func).Call(args)
			s += result[0].Interface().(string) + "\n"
		} else {
			found := false
			for _, t := range htmlTags {
				if t == x.Name {
					found = true
				}
			}
			if !found {
				panic(fmt.Errorf("Comp not found %s", x.Name))
			}
			for _, c := range x.Children {
				ctx["_space"] = space + "  "
				s += render(c, ctx) + "\n"
			}
		}
	}
	s += space + "</" + x.Name + ">"
	return s
}

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
