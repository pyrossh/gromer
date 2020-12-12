package app

import (
	"io"

	"github.com/pyros2097/wapp/errors"
	"github.com/pyros2097/wapp/js"
)

// UI is the interface that describes a user interface element such as
// components and HTML elements.
type UI interface {
	// JSValue returns the javascript value linked to the element.
	JSValue() js.Value

	// Reports whether the element is mounted.
	Mounted() bool

	name() string
	self() UI
	setSelf(UI)
	attributes() map[string]string
	eventHandlers() map[string]js.EventHandler
	parent() UI
	setParent(UI)
	children() []UI
	mount() error
	dismount()
	update(UI) error
}

func trackMousePosition(e js.Event) {
	x := e.Get("clientX")
	if !x.Truthy() {
		return
	}

	y := e.Get("clientY")
	if !y.Truthy() {
		return
	}

	js.Window.SetCursorPosition(x.Int(), y.Int())
}

func isErrReplace(err error) bool {
	_, replace := errors.Tag(err, "replace")
	return replace
}

func mount(n UI) error {
	n.setSelf(n)
	return n.mount()
}

func dismount(n UI) {
	n.dismount()
	n.setSelf(nil)
}

func update(a, b UI) error {
	a.setSelf(a)
	b.setSelf(b)
	return a.update(b)
}

type WritableNode interface {
	Html(w io.Writer)
	HtmlWithIndent(w io.Writer, indent int)
}
