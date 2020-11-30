package app

import (
	"context"
)

type baseAttribute struct {
	UI
}

func (c baseAttribute) Kind() Kind {
	return Attribute
}

func (c baseAttribute) JSValue() Value {
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

func (c baseAttribute) context() context.Context {
	return nil
}

func (c baseAttribute) attributes() map[string]string {
	return nil
}

func (c baseAttribute) eventHandlers() map[string]eventHandler {
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
