package app

import (
	"io"
	"reflect"
	"strconv"
)

type Element struct {
	tag         string
	attrs       map[string]string
	body        []*Element
	selfClosing bool
	text        string
}

func NewElement(tag string, selfClosing bool, uis ...interface{}) *Element {
	return MergeAttributes(&Element{tag: tag, selfClosing: selfClosing}, uis...)
}

func (e *Element) setAttr(k string, v string) {
	if e.attrs == nil {
		e.attrs = make(map[string]string)
	}

	switch k {
	case "style", "allow":
		s := e.attrs[k] + v + ";"
		e.attrs[k] = s
		return

	case "class":
		s := e.attrs[k]
		if s != "" {
			s += " "
		}
		s += v
		e.attrs[k] = s
		return
	}
	if v == "false" {
		delete(e.attrs, k)
		return
	} else if v == "true" {
		e.attrs[k] = ""
	} else {
		e.attrs[k] = v
	}
}

func (e *Element) WriteHtml(w io.Writer) {
	e.writeHtmlIndent(w, 0)
}

func (e *Element) writeHtmlIndent(w io.Writer, indent int) {
	e.writeIndent(w, indent)
	if e.tag == "html" {
		w.Write([]byte("<!DOCTYPE html>\n"))
	}
	if e.tag == "text" {
		e.writeIndent(w, indent)
		w.Write([]byte(e.text))
		return
	}
	w.Write([]byte("<"))
	w.Write([]byte(e.tag))

	for k, v := range e.attrs {
		w.Write([]byte(" "))
		w.Write([]byte(k))

		if v != "" {
			w.Write([]byte(`="`))
			w.Write([]byte(v))
			w.Write([]byte(`"`))
		}
	}

	w.Write([]byte(">"))

	if e.selfClosing {
		return
	}

	for _, c := range e.body {
		w.Write([]byte("\n"))
		if c != nil {
			c.writeHtmlIndent(w, indent+1)
		}
	}

	if len(e.body) != 0 {
		w.Write([]byte("\n"))
		e.writeIndent(w, indent)
	}

	w.Write([]byte("</"))
	w.Write([]byte(e.tag))
	w.Write([]byte(">"))
}

func (e *Element) writeIndent(w io.Writer, indent int) {
	for i := 0; i < indent*4; i++ {
		w.Write([]byte(" "))
	}
}

func Html(elems ...*Element) *Element {
	return &Element{tag: "html", body: elems}
}

func Head(elems ...*Element) *Element {
	basic := []*Element{
		&Element{tag: "meta", selfClosing: true, attrs: map[string]string{"charset": "UTF-8"}},
		&Element{tag: "meta", selfClosing: true, attrs: map[string]string{"http-equiv": "Content-Type", "content": "text/html;charset=utf-8"}},
		&Element{tag: "meta", selfClosing: true, attrs: map[string]string{"http-equiv": "encoding", "content": "utf-8"}},
	}
	return &Element{tag: "head", body: append(basic, elems...)}
}

func Body(elems ...*Element) *Element {
	return &Element{tag: "body", body: elems}
}

func Title(v string) *Element {
	return &Element{
		tag:  "title",
		body: []*Element{Text(v)},
	}
}
func Text(v string) *Element {
	return &Element{
		tag:  "text",
		text: v,
	}
}

func Meta(name, content string) *Element {
	return &Element{
		tag:         "meta",
		selfClosing: true,
		attrs: map[string]string{
			"name":    name,
			"content": content,
		},
	}
}

func Link(rel, href string) *Element {
	return &Element{
		tag:         "link",
		selfClosing: true,
		attrs: map[string]string{
			"rel":  rel,
			"href": href,
		},
	}
}

func Script(uis ...interface{}) *Element {
	return NewElement("script", false, uis...)
}

func Div(uis ...interface{}) *Element {
	return NewElement("div", false, uis...)
}

func A(uis ...interface{}) *Element {
	return NewElement("a", false, uis...)
}

func P(uis ...interface{}) *Element {
	return NewElement("p", false, uis...)
}

func H1(uis ...interface{}) *Element {
	return NewElement("h1", false, uis...)
}
func H2(uis ...interface{}) *Element {
	return NewElement("h2", false, uis...)
}
func H3(uis ...interface{}) *Element {
	return NewElement("h3", false, uis...)
}
func H4(uis ...interface{}) *Element {
	return NewElement("h4", false, uis...)
}
func H5(uis ...interface{}) *Element {
	return NewElement("h5", false, uis...)
}
func H6(uis ...interface{}) *Element {
	return NewElement("h6", false, uis...)
}

func Span(uis ...interface{}) *Element {
	return NewElement("span", false, uis...)
}

func Input(uis ...interface{}) *Element {
	return NewElement("input", false, uis...)
}

func Image(uis ...interface{}) *Element {
	return NewElement("image", true, uis...)
}

func Button(uis ...interface{}) *Element {
	return NewElement("button", false, uis...)
}

func Svg(uis ...interface{}) *Element {
	return NewElement("svg", false, uis...)
}

func SvgText(uis ...interface{}) *Element {
	return NewElement("text", false, uis...)
}

func Ul(uis ...interface{}) *Element {
	return NewElement("ul", false, uis...)
}

func Li(uis ...interface{}) *Element {
	return NewElement("li", false, uis...)
}

func Row(uis ...interface{}) *Element {
	return Div(append([]interface{}{Css("flex flex-row justify-center items-center")}, uis...)...)
}

func Col(uis ...interface{}) *Element {
	return Div(append([]interface{}{Css("flex flex-col justify-center items-center")}, uis...)...)
}

func If(expr bool, a *Element) *Element {
	if expr {
		return a
	}
	return nil
}

func IfElse(expr bool, a *Element, b *Element) *Element {
	if expr {
		return a
	}
	return b
}

func Map(source interface{}, f func(i int) *Element) []*Element {
	src := reflect.ValueOf(source)
	if src.Kind() != reflect.Slice {
		panic("range loop source is not a slice: " + src.Type().String())
	}
	body := make([]*Element, 0, src.Len())
	for i := 0; i < src.Len(); i++ {
		body = append(body, f(i))
	}
	return body
}

func Map2(source interface{}, f func(v interface{}, i int) *Element) []*Element {
	src := reflect.ValueOf(source)
	if src.Kind() != reflect.Slice {
		panic("range loop source is not a slice: " + src.Type().String())
	}
	body := make([]*Element, 0, src.Len())
	for i := 0; i < src.Len(); i++ {
		body = append(body, f(src.Index(i), i))
	}
	return body
}

type Attribute struct {
	Key   string
	Value string
}

func ID(v string) Attribute {
	return Attribute{"id", v}
}

func Style(v string) Attribute {
	return Attribute{"style", v}
}

func Accept(v string) Attribute {
	return Attribute{"accept", v}
}

func AutoComplete(v bool) Attribute {
	return Attribute{"autocomplete", strconv.FormatBool(v)}
}

func Checked(v bool) Attribute {
	return Attribute{"checked", strconv.FormatBool(v)}
}

func Disabled(v bool) Attribute {
	return Attribute{"disabled", strconv.FormatBool(v)}
}

func Name(v string) Attribute {
	return Attribute{"name", v}
}

func Type(v string) Attribute {
	return Attribute{"type", v}
}

func Value(v string) Attribute {
	return Attribute{"value", v}
}

func Placeholder(v string) Attribute {
	return Attribute{"placeholder", v}
}

func Src(v string) Attribute {
	return Attribute{"src", v}
}

func Defer() Attribute {
	return Attribute{"defer", "true"}
}

func ViewBox(v string) Attribute {
	return Attribute{"viewBox", v}
}

func X(v string) Attribute {
	return Attribute{"x", v}
}

func Y(v string) Attribute {
	return Attribute{"y", v}
}

func Href(v string) Attribute {
	return Attribute{"href", v}
}

func Target(v string) Attribute {
	return Attribute{"target", v}
}

func Rel(v string) Attribute {
	return Attribute{"rel", v}
}

func Css(v string) Attribute {
	return Attribute{"class", v}
}

func XData(v string) Attribute {
	return Attribute{"x-data", v}
}

func XText(v string) Attribute {
	return Attribute{"x-text", v}
}

func MergeAttributes(parent *Element, uis ...interface{}) *Element {
	elems := []*Element{}
	for _, v := range uis {
		switch c := v.(type) {
		case Attribute:
			parent.setAttr(c.Key, c.Value)
		case *Element:
			elems = append(elems, c)
		case nil:
			// dont need to add nil items
		default:
			// fmt.Printf("%v\n", v)
			panic("unknown type in render")
		}
	}
	if !parent.selfClosing {
		parent.body = elems
	}
	return parent
}
