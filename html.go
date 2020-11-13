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
	return &elem{tag: "head", body: elems}
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

func Div(elems ...UI) *elem {
	return &elem{tag: "div", body: elems}
}

func Row(elems ...UI) *elem {
	return &elem{tag: "div", body: elems, attrs: map[string]string{
		"style": "display: flex;flex: 1;flex-direction: row;align-items: center;justify-content: center;",
	}}
}

func Col(elems ...UI) *elem {
	return &elem{tag: "div", body: elems, attrs: map[string]string{
		"style": "display: flex;flex: 1;flex-direction: column;align-items: center;justify-content: center;",
	}}
}
