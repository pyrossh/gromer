package components

import (
	. "github.com/pyros2097/wapp"
)

func Page(elem *Element) *Element {
	return Html(
		Head(
			Title("123"),
			Meta("description", "123"),
			Meta("author", "123"),
			Meta("keywords", "123"),
			Meta("viewport", "width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover"),
			Link("icon", "/assets/icon.png"),
			Link("apple-touch-icon", "/assets/icon.png"),
			Link("stylesheet", "/assets/styles.css"),
			Script(Src("/assets/alpine.js"), Defer()),
		),
		Body(elem),
	)
}
