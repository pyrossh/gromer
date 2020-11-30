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

// func (e *elem) OnBlur(h EventHandler) *elem {
// 	e.setEventHandler("blur", h)
// 	return e
// }

// func (e *elem) OnChange(h EventHandler) *elem {
// 	e.setEventHandler("change", h)
// 	return e
// }

// func (e *elem) OnFocus(h EventHandler) *elem {
// 	e.setEventHandler("focus", h)
// 	return e
// }

// func (e *elem) OnInput(h EventHandler) *elem {
// 	e.setEventHandler("input", h)
// 	return e
// }

func mergeAttributes(parent *elem, uis ...UI) {
	elems := make([]UI, 0, len(uis))
	for _, v := range uis {
		if v.Kind() == Attribute {
			switch c := v.(type) {
			case CssAttribute:
				if vv, ok := parent.attrs["classes"]; ok {
					parent.setAttr("class", vv+" "+c.classes)
				} else {
					parent.setAttr("class", c.classes)
				}
			case OnClickAttribute:
				parent.setEventHandler("click", func(e Event) {
					c.cb()
				})
			}
		} else {
			elems = append(elems, v)
		}

	}
	parent.setBody(elems...)
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
