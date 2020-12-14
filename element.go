package app

import (
	"io"
	"strconv"

	"github.com/pyros2097/wapp/js"
)

type elem struct {
	attrs       map[string]string
	body        []UI
	events      map[string]js.EventHandler
	jsvalue     js.Value
	parentElem  UI
	selfClosing bool
	tag         string
	this        UI
}

func (e *elem) JSValue() js.Value {
	return e.jsvalue
}

func (e *elem) Mounted() bool {
	return e.self() != nil &&
		e.jsvalue != nil
}

func (e *elem) name() string {
	return e.tag
}

func (e *elem) self() UI {
	return e.this
}

func (e *elem) setSelf(n UI) {
	e.this = n
}

func (e *elem) attributes() map[string]string {
	return e.attrs
}

func (e *elem) eventHandlers() map[string]js.EventHandler {
	return e.events
}

func (e *elem) parent() UI {
	return e.parentElem
}

func (e *elem) setParent(p UI) {
	e.parentElem = p
}

func (e *elem) children() []UI {
	return e.body
}

func (e *elem) mount() error {
	if e.Mounted() {
		panic("mounting elem failed already mounted " + e.name())
	}

	v := js.Window.Get("document").Call("createElement", e.tag)
	if !v.Truthy() {
		panic("mounting component failed create javascript node returned nil " + e.name())
	}
	e.jsvalue = v

	for k, v := range e.attrs {
		e.setJsAttr(k, v)
	}

	for k, v := range e.events {
		e.setJsEventHandler(k, v)
	}

	for _, c := range e.children() {
		if err := e.appendChild(c, true); err != nil {
			panic("mounting component failed appendChild " + e.name())
		}
	}

	return nil
}

func (e *elem) dismount() {
	for _, c := range e.children() {
		dismount(c)
	}

	for k, v := range e.events {
		e.delJsEventHandler(k, v)
	}

	e.jsvalue = nil
}

func (e *elem) update(n UI) error {
	if !e.Mounted() {
		return nil
	}

	if n.name() != e.name() {
		panic("updating element failed replace different element type current-name: " + e.name() + " updated-name: " + n.name())
	}

	e.updateAttrs(n.attributes())
	e.updateEventHandler(n.eventHandlers())

	achildren := e.children()
	bchildren := n.children()
	i := 0

	// Update children:
	for len(achildren) != 0 && len(bchildren) != 0 {
		a := achildren[0]
		b := bchildren[0]

		err := update(a, b)
		if isErrReplace(err) {
			err = e.replaceChildAt(i, b)
		}

		if err != nil {
			panic("updating element failed name: " + e.name())
		}

		achildren = achildren[1:]
		bchildren = bchildren[1:]
		i++
	}

	// Remove children:
	for len(achildren) != 0 {
		if err := e.removeChildAt(i); err != nil {
			panic("updating element failed name: " + e.name())
		}

		achildren = achildren[1:]
	}

	// Add children:
	for len(bchildren) != 0 {
		c := bchildren[0]

		if err := e.appendChild(c, false); err != nil {
			panic("updating element failed name: " + e.name())
		}

		bchildren = bchildren[1:]
	}

	return nil
}

func (e *elem) appendChild(c UI, onlyJsValue bool) error {
	if err := mount(c); err != nil {
		panic("appending child failed child-name: " + c.name() + " name: " + e.name())
	}

	if !onlyJsValue {
		e.body = append(e.body, c)
	}

	c.setParent(e.self())
	e.JSValue().Call("appendChild", c)
	return nil
}

func (e *elem) replaceChildAt(idx int, new UI) error {
	old := e.body[idx]

	if err := mount(new); err != nil {
		panic("replacing child failed name: " + e.name() + " old-name: " + old.name() + "  new-name: " + new.name())
	}

	e.body[idx] = new
	new.setParent(e.self())
	e.JSValue().Call("replaceChild", new, old)

	dismount(old)
	return nil
}

func (e *elem) removeChildAt(idx int) error {
	body := e.body
	if idx < 0 || idx >= len(body) {
		panic("removing child failed index out of range name: " + e.name() + " index: " + strconv.Itoa(idx))
	}

	c := body[idx]

	copy(body[idx:], body[idx+1:])
	body[len(body)-1] = nil
	body = body[:len(body)-1]
	e.body = body

	e.JSValue().Call("removeChild", c)
	dismount(c)
	return nil
}

func (e *elem) updateAttrs(attrs map[string]string) {
	for k := range e.attrs {
		if _, exists := attrs[k]; !exists {
			e.delAttr(k)
		}
	}

	if e.attrs == nil && len(attrs) != 0 {
		e.attrs = make(map[string]string, len(attrs))
	}

	for k, v := range attrs {
		if curval, exists := e.attrs[k]; !exists || curval != v {
			e.attrs[k] = v
			e.setJsAttr(k, v)
		}
	}
}

func (e *elem) setAttr(k string, v string) {
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

func (e *elem) setJsAttr(k, v string) {
	e.JSValue().Call("setAttribute", k, v)
}

func (e *elem) delAttr(k string) {
	e.JSValue().Call("removeAttribute", k)
	delete(e.attrs, k)
}

func (e *elem) updateEventHandler(handlers map[string]js.EventHandler) {
	for k, current := range e.events {
		if _, exists := handlers[k]; !exists {
			e.delJsEventHandler(k, current)
		}
	}

	if e.events == nil && len(handlers) != 0 {
		e.events = make(map[string]js.EventHandler, len(handlers))
	}

	for k, new := range handlers {
		if current, exists := e.events[k]; !current.Equal(new) {
			if exists {
				e.delJsEventHandler(k, current)
			}

			e.events[k] = new
			e.setJsEventHandler(k, new)
		}
	}
}

func (e *elem) setEventHandler(k string, h js.EventHandlerFunc) {
	if e.events == nil {
		e.events = make(map[string]js.EventHandler)
	}

	e.events[k] = js.NewEventHandler(k, h)
}

func (e *elem) setJsEventHandler(k string, h js.EventHandler) {
	jshandler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		dispatch(func() {
			if !e.self().Mounted() {
				return
			}
			e := js.Event{
				Src:   this,
				Value: args[0],
			}
			trackMousePosition(e)
			h.Value(e)
		})
		return nil
	})
	h.JSvalue = jshandler
	e.events[k] = h
	e.JSValue().Call("addEventListener", k, jshandler)
}

func (e *elem) delJsEventHandler(k string, h js.EventHandler) {
	e.JSValue().Call("removeEventListener", k, h.JSvalue)
	h.JSvalue.Release()
	delete(e.events, k)
}

func (e *elem) setBody(body []UI) {
	if e.selfClosing {
		panic("setting html element body failed: self closing element can't have children" + e.name())
	}
	e.body = body
}

func (e *elem) Html(w io.Writer) {
	e.HtmlWithIndent(w, 0)
}

func (e *elem) HtmlWithIndent(w io.Writer, indent int) {
	writeIndent(w, indent)
	w.Write(stob("<"))
	w.Write(stob(e.tag))

	for k, v := range e.attrs {
		w.Write(stob(" "))
		w.Write(stob(k))

		if v != "" {
			w.Write(stob(`="`))
			w.Write(stob(v))
			w.Write(stob(`"`))
		}
	}

	w.Write(stob(">"))

	if e.selfClosing {
		return
	}

	for _, c := range e.body {
		w.Write(ln())
		c.(WritableNode).HtmlWithIndent(w, indent+1)
	}

	if len(e.body) != 0 {
		w.Write(ln())
		writeIndent(w, indent)
	}

	w.Write(stob("</"))
	w.Write(stob(e.tag))
	w.Write(stob(">"))
}

type text struct {
	jsvalue    js.Value
	parentElem UI
	value      string
}

// Text creates a simple text element.
func Text(v string) UI {
	return &text{value: v}
}

func (t *text) JSValue() js.Value {
	return t.jsvalue
}

func (t *text) Mounted() bool {
	return t.jsvalue != nil
}

func (t *text) name() string {
	return "text"
}

func (t *text) self() UI {
	return t
}

func (t *text) setSelf(n UI) {
}

func (t *text) attributes() map[string]string {
	return nil
}

func (t *text) eventHandlers() map[string]js.EventHandler {
	return nil
}

func (t *text) parent() UI {
	return t.parentElem
}

func (t *text) setParent(p UI) {
	t.parentElem = p
}

func (t *text) children() []UI {
	return nil
}

func (t *text) mount() error {
	if t.Mounted() {
		panic("mounting text element failed already mounted" + t.name() + " " + t.value)
	}

	t.jsvalue = js.Window.
		Get("document").
		Call("createTextNode", t.value)

	return nil
}

func (t *text) dismount() {
	t.jsvalue = nil
}

func (t *text) update(n UI) error {
	if !t.Mounted() {
		return nil
	}

	o, isText := n.(*text)
	if !isText {
		panic("updating ui element failed replace different element type current-name: " + t.name() + " updated-name: " + n.name())
	}

	if t.value != o.value {
		t.value = o.value
		t.jsvalue.Set("nodeValue", o.value)
	}

	return nil
}

func (t *text) Html(w io.Writer) {
	t.HtmlWithIndent(w, 0)
}

func (t *text) HtmlWithIndent(w io.Writer, indent int) {
	writeIndent(w, indent)
	// html.EscapeString(
	w.Write(stob(t.value))
}
