package app

func Html(elems ...UI) *elem {
	return &elem{tag: "html", body: elems}
}

func Head(elems ...UI) *elem {
	basic := []UI{
		&elem{tag: "meta", selfClosing: true, attrs: map[string]string{"charset": "UTF-8"}},
		&elem{tag: "meta", selfClosing: true, attrs: map[string]string{"http-equiv": "Content-Type", "content": "text/html;charset=utf-8"}},
		&elem{tag: "meta", selfClosing: true, attrs: map[string]string{"http-equiv": "encoding", "content": "utf-8"}},
	}
	return &elem{tag: "head", body: append(basic, elems...)}
}

func Body(elems ...UI) *elem {
	return &elem{tag: "body", body: elems}
}

func Title(v string) *elem {
	return &elem{tag: "title", body: []UI{Text(v)}}
}

func Meta(name, content string) *elem {
	e := &elem{
		tag:         "meta",
		selfClosing: true,
	}
	e.setAttr("name", name)
	e.setAttr("content", content)
	return e
}

func Link(rel, href string) *elem {
	e := &elem{
		tag:         "link",
		selfClosing: true,
	}
	e.setAttr("rel", rel)
	e.setAttr("href", href)
	return e
}

func Script(str string) *elem {
	return &elem{
		tag:  "script",
		body: []UI{Text(str)},
	}
}

func Div(uis ...UI) *elem {
	e := &elem{tag: "div"}
	mergeAttributes(e, uis...)
	return e
}

func Row(uis ...UI) UI {
	return Div(append([]UI{Css("flex flex-row justify-center align-items-center")}, uis...)...)
}

func Col(uis ...UI) UI {
	return Div(append([]UI{Css("flex flex-col justify-center align-items-center")}, uis...)...)
}
