package gsx

import (
	"fmt"
	"io"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/samber/lo"
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
		Name    string
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
		Name:    name,
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
			t := funcType.In(i + 1)
			v, _ := lo.Find(tag.Attributes, func(a *Attribute) bool {
				return a.Key == arg
			})
			var data interface{}
			if v.Value.Ref != nil {
				data = getRefValue(c, *v.Value.Ref)
			} else if v.Value.Str != nil {
				data = *v.Value.Str
			}
			switch t.Kind() {
			case reflect.Int:
				var value int
				if v, ok := data.(int); ok {
					value = v
				} else {
					s, ok := data.(string)
					if !ok {
						panic(eris.Errorf("expected component %s: prop %s to be of type string but got %+v ", comp.Name, arg, data))
					}
					value, _ = strconv.Atoi(s)
				}
				c.Set(arg, value)
				args = append(args, reflect.ValueOf(value))
			case reflect.Bool:
				var value bool
				if v, ok := data.(bool); ok {
					value = v
				} else {
					s, ok := data.(string)
					if !ok {
						panic(eris.Errorf("expected component %s: prop %s to be of type string but got %+v ", comp.Name, arg, data))
					}
					value, _ = strconv.ParseBool(s)
				}
				c.Set(arg, value)
				args = append(args, reflect.ValueOf(value))
			default:
				c.Set(arg, data)
				args = append(args, reflect.ValueOf(data))
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
		w.Write([]byte(`</head><body _="on htmx:error(errorInfo) put errorInfo.xhr.response into #error">`))
	}
	out := RenderString(tags)
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
	case []*Tag:
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

func removeBrackets(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "{", ""), "}", "")
}

func removeQuotes(s string) string {
	return strings.ReplaceAll(s, `"`, "")
}

func substituteString(c *Context, v string) string {
	found := refRegex.FindString(v)
	if found != "" {
		varValue := fmt.Sprintf("%v", getRefValue(c, removeBrackets(found)))
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
		if tag.Text.Str == nil && tag.Text.Ref != nil {
			value := getRefValue(c, *tag.Text.Ref)
			children, ok := value.([]*Tag)
			if ok {
				tag.Name = "fragment"
				tag.Children = children
			} else {
				sValue := fmt.Sprintf("%+v", value)
				tag.Text.Str = &sValue
			}
		} else if loop := tag.Text.For; loop != nil {
			tag.Name = "fragment"
			data := getRefValue(c, loop.Reference)
			statement := loop.Statements[0].ReturnStatement
			switch reflect.TypeOf(data).Kind() {
			case reflect.Slice:
				v := reflect.ValueOf(data)
				for i := 0; i < v.Len(); i++ {
					compContext := c.Clone(tag.Name)
					compContext.data[loop.Index] = i
					compContext.data[loop.Key] = v.Index(i).Interface()
					newTags := populate(compContext, cloneTags(statement.Tags))
					tag.Children = append(tag.Children, newTags...)
				}
			}
		}
	} else {
		if comp, ok := compMap[tag.Name]; ok {
			tag.Name = "fragment"
			if tag.SelfClosing {
				tag.SelfClosing = false
			}
			compContext := c.Clone(comp.Name)
			nodes := comp.Render(compContext, tag)
			populate(compContext, tag.Children)
			compContext.Set("children", tag.Children)
			populate(compContext, nodes)
			tag.Children = nodes
		} else {
			for _, a := range tag.Attributes {
				if a.Value.Str != nil {
					if strings.Contains(*a.Value.Str, "{") {
						subs := substituteString(c, removeQuotes(*a.Value.Str))
						a.Value = &Literal{Str: &subs}
					} else {
						*a.Value.Str = removeQuotes(*a.Value.Str)
					}
				} else if a.Value.Ref != nil {
					value := getRefValue(c, *a.Value.Ref)
					subs := fmt.Sprintf("%v", value)
					a.Value = &Literal{Str: &subs}
				} else if a.Key == "class" && a.Value.KV != nil {
					classes := []string{}
					for _, a := range a.Value.KV {
						varValue := getRefValue(c, a.Value)
						if varValue.(bool) {
							classes = append(classes, removeQuotes(a.Key))
						}
					}
					result := strings.Join(classes, " ")
					a.Value.Str = &result
				}
			}
			populate(c, tag.Children)
		}
	}
}
