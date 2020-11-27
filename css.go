package app

import (
	"context"

	"github.com/pyros2097/wapp/errors"
)

type CSSClass struct {
	UI
	classes string
}

func Css(d string) UI {
	return CSSClass{classes: d}
}

func (c CSSClass) Kind() Kind {
	return Attribute
}

func (c CSSClass) JSValue() Value {
	return nil
}

func (c CSSClass) Mounted() bool {
	return false
}

func (c CSSClass) name() string {
	return "css"
}

func (c CSSClass) self() UI {
	return c
}

func (c CSSClass) setSelf(UI) {
}

func (c CSSClass) context() context.Context {
	return nil
}

func (c CSSClass) attributes() map[string]string {
	return nil
}

func (c CSSClass) eventHandlers() map[string]eventHandler {
	return nil
}

func (c CSSClass) parent() UI {
	return nil
}

func (c CSSClass) setParent(UI) {
}

func (c CSSClass) children() []UI {
	return nil
}

func (c CSSClass) mount() error {
	return errors.New("condition is not mountable").
		Tag("name", c.name()).
		Tag("kind", c.Kind())
}

func (c CSSClass) dismount() {
}

func (c CSSClass) update(UI) error {
	return errors.New("condition cannot be updated").
		Tag("name", c.name()).
		Tag("kind", c.Kind())
}
