package gsx

import (
	"fmt"
	"io"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var (
	htmlElements = []string{"a", "abbr", "acronym", "address", "applet", "area", "article", "aside", "audio", "b", "base", "basefont", "bb", "bdo", "big", "blockquote", "body", "br /", "button", "canvas", "caption", "center", "cite", "code", "col", "colgroup", "command", "datagrid", "datalist", "dd", "del", "details", "dfn", "dialog", "dir", "div", "dl", "dt", "em", "embed", "eventsource", "fieldset", "figcaption", "figure", "font", "footer", "form", "frame", "frameset", "h1 to <h6>", "head", "header", "hgroup", "hr /", "html", "i", "iframe", "img", "input", "ins", "isindex", "kbd", "keygen", "label", "legend", "li", "link", "map", "mark", "menu", "meta", "meter", "nav", "noframes", "noscript", "object", "ol", "optgroup", "option", "output", "p", "param", "pre", "progress", "q", "rp", "rt", "ruby", "s", "samp", "script", "section", "select", "small", "source", "span", "strike", "strong", "style", "sub", "sup", "table", "tbody", "td", "textarea", "tfoot", "th", "thead", "time", "title", "tr", "track", "tt", "u", "ul", "var", "video", "wbr"}
	voidElements = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}
	compMap      = map[string]ComponentFunc{}
	funcMap      = map[string]interface{}{}
	classesMap   = map[string]M{}
	refRegex     = regexp.MustCompile(`{(.*?)}`)
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
)

func RegisterComponent(f interface{}, classes M, args ...string) {
	name := getFunctionName(f)
	compMap[name] = ComponentFunc{
		Func:    f,
		Args:    args,
		Classes: classes,
	}
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
	funcMap[name] = f
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func (comp ComponentFunc) Render(c *Context, tag *Tag) []*Tag {
	args := []reflect.Value{reflect.ValueOf(c)}
	funcType := reflect.TypeOf(comp.Func)
	for i, arg := range comp.Args {
		if v, ok := c.data[arg]; ok {
			args = append(args, reflect.ValueOf(v))
		} else {
			v := findAttribute(tag.Attributes, arg)
			t := funcType.In(i + 1)
			switch t.Kind() {
			case reflect.Int:
				value, _ := strconv.Atoi(*v.Value.Str)
				args = append(args, reflect.ValueOf(value))
			case reflect.Bool:
				value, _ := strconv.ParseBool(*v.Value.Str)
				args = append(args, reflect.ValueOf(value))
			default:
				args = append(args, reflect.ValueOf(v))
			}
		}
	}
	result := reflect.ValueOf(comp.Func).Call(args)
	tags := result[0].Interface().([]*Tag)
	return populate(c, tags)
}

func Write(c *Context, w io.Writer, tags []*Tag) {
	if c.hx == nil {
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
	out := renderString(tags)
	w.Write([]byte(out))
	if c.hx == nil {
		w.Write([]byte(`</body></html>`))
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

func substituteString(c *Context, v string) string {
	found := refRegex.FindString(v)
	if found != "" {
		varValue := fmt.Sprintf("%v", getRefValue(c, found))
		return strings.ReplaceAll(v, found, varValue)
	}
	return v
}

func populate(c *Context, tags []*Tag) []*Tag {
	for _, t := range tags {
		populateTag(c, t)
	}
	return tags
}

func populateTag(c *Context, tag *Tag) {
	if tag.Name == "" {
		if tag.Text.Ref != nil && *tag.Text.Ref != "children" {
			*tag.Text.Ref = substituteString(c, *tag.Text.Ref)
		}
	} else {
		for _, a := range tag.Attributes {
			if a.Key == "x-for" {
				arr := strings.Split(*a.Value.Str, " in ")
				// ctxItemKey := arr[0]
				ctxKey := arr[1]
				data := c.data[ctxKey]
				switch reflect.TypeOf(data).Kind() {
				case reflect.Slice:
					v := reflect.ValueOf(data)
					for i := 0; i < v.Len(); i++ {
						// ctx["_space"] = space + "  "
						// ctx[ctxName] = v.Index(i).Interface()
						// s += render(x.Children[0], ctx) + "\n"

						// compCtx := &Context{
						// 	Context: c.Context,
						// 	data: map[string]interface{}{
						// 		ctxItemKey: v.Index(i).Interface(),
						// 	},
						// }
						// tag.Children
						// if comp, ok := compMap[itemChild.Data]; ok {
						// 	newNode := populateComponent(compCtx, comp, itemChild, false)
						// 	n.AppendChild(newNode)
						// } else {
						// 	n.AppendChild(itemChild)
						// 	populate(compCtx, itemChild)
						// }
					}
				}
			} else if a.Value.Ref != nil {
				if a.Key == "class" {
					// 	classes := []string{}
					// 	kvstrings := strings.Split(strings.TrimSpace(at.Val), ",")
					// 	for _, kv := range kvstrings {
					// 		kvarray := strings.Split(kv, ":")
					// 		k := strings.TrimSpace(kvarray[0])
					// 		v := strings.TrimSpace(kvarray[1])
					// 		varValue := getRefValue(c, v)
					// 		if varValue.(bool) {
					// 			classes = append(classes, k)
					// 		}
					// 	}
					// 	n.Attr[i] = html.Attribute{
					// 		Namespace: at.Namespace,
					// 		Key:       at.Key,
					// 		Val:       strings.Join(classes, " "),
					// 	}
				} else {
					subs := substituteString(c, *a.Value.Ref)
					a.Value = &Literal{Str: &subs}
				}
			}
		}
		if comp, ok := compMap[tag.Name]; ok {
			compContext := c.Clone(tag.Name)
			nodes := comp.Render(compContext, tag)
			// TODO: check if tag as Children already and {children} in defined in component
			tag.Children = nodes
		}
		populate(c, tag.Children)
	}
}
