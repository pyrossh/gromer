package app

import (
	"io"
	"reflect"

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

// FilterUIElems returns a filtered version of the given UI elements where
// selector elements such as If and Range are interpreted and removed. It also
// remove nil elements.
//
// It should be used only when implementing components that can accept content
// with variadic arguments like HTML elements Body method.
func FilterUIElems(uis ...interface{}) []UI {
	if len(uis) == 0 {
		return nil
	}

	elems := make([]UI, 0, len(uis))

	for _, n := range uis {
		if v := reflect.ValueOf(n); n == nil ||
			v.Kind() == reflect.Ptr && v.IsNil() {
			continue
		}
		switch c := n.(type) {
		case RangeLoop:
			elems = append(elems, c.body...)
		case Condition:
			elems = append(elems, c.body...)
		default:
			elems = append(elems, c.(UI))
		}
	}

	return elems
}

func makeJsEventHandler(src UI, h js.EventHandlerFunc) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		dispatch(func() {
			if !src.Mounted() {
				return
			}
			e := js.Event{
				Value: args[0],
			}
			trackMousePosition(e)
			h(e)
		})

		return nil
	})
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
