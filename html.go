package app

import "reflect"

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
	return NewElement("h1", false, uis...)
}
func H3(uis ...interface{}) *Element {
	return NewElement("h1", false, uis...)
}
func H4(uis ...interface{}) *Element {
	return NewElement("h1", false, uis...)
}
func H5(uis ...interface{}) *Element {
	return NewElement("h1", false, uis...)
}
func H6(uis ...interface{}) *Element {
	return NewElement("h1", false, uis...)
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
