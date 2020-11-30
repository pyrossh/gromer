package app

import (
	"github.com/pyros2097/wapp/js"
)

type baseAttribute struct {
	UI
}

func (c baseAttribute) Kind() Kind {
	return Attribute
}

func (c baseAttribute) JSValue() js.Value {
	return nil
}

func (c baseAttribute) Mounted() bool {
	return false
}

func (c baseAttribute) name() string {
	return "css"
}

func (c baseAttribute) self() UI {
	return c
}

func (c baseAttribute) setSelf(UI) {
}

func (c baseAttribute) attributes() map[string]string {
	return nil
}

func (c baseAttribute) eventHandlers() map[string]js.EventHandler {
	return nil
}

func (c baseAttribute) parent() UI {
	return nil
}

func (c baseAttribute) setParent(UI) {
}

func (c baseAttribute) children() []UI {
	return nil
}

func (c baseAttribute) mount() error {
	panic("cant mount attributes")
}

func (c baseAttribute) dismount() {
}

func (c baseAttribute) update(UI) error {
	panic("cant update attributes")
}

type CssAttribute struct {
	baseAttribute
	classes string
}

func Css(d string) UI {
	return CssAttribute{classes: d}
}

type OnClickAttribute struct {
	baseAttribute
	cb func()
}

func OnClick(cb func()) UI {
	return OnClickAttribute{cb: cb}
}

type OnChangeAttribute struct {
	baseAttribute
	cb js.EventHandlerFunc
}

func OnChange(cb js.EventHandlerFunc) UI {
	return OnChangeAttribute{cb: cb}
}

type OnInputAttribute struct {
	baseAttribute
	cb js.EventHandlerFunc
}

func OnInput(cb js.EventHandlerFunc) UI {
	return OnInputAttribute{cb: cb}
}

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
				parent.setEventHandler("click", func(e js.Event) {
					c.cb()
				})
			case OnChangeAttribute:
				parent.setEventHandler("change", c.cb)
			case OnInputAttribute:
				parent.setEventHandler("input", c.cb)
			}
		} else {
			elems = append(elems, v)
		}

	}
	parent.setBody(elems...)
}
