package app

import (
	"context"
	"net/url"
	"reflect"
	"sort"

	"github.com/pyros2097/wapp/errors"
)

// RangeLoop represents a control structure that iterates within a slice, an
// array or a map.
type RangeLoop interface {
	UI

	// Slice sets the loop content by repeating the given function for the
	// number of elements in the source.
	//
	// It panics if the range source is not a slice or an array.
	Slice(f func(int) UI) RangeLoop

	// Map sets the loop content by repeating the given function for the number
	// of elements in the source. Elements are ordered by keys.
	//
	// It panics if the range source is not a map or if map keys are not strings.
	Map(f func(string) UI) RangeLoop
}

// Range returns a range loop that iterates within the given source. Source must
// be a slice, an array or a map with strings as keys.
func Range(src interface{}) RangeLoop {
	return rangeLoop{source: src}
}

type rangeLoop struct {
	body   []UI
	source interface{}
}

func (r rangeLoop) Slice(f func(int) UI) RangeLoop {
	src := reflect.ValueOf(r.source)
	if src.Kind() != reflect.Slice && src.Kind() != reflect.Array {
		panic(errors.New("range loop source is not a slice or array").
			Tag("src-type", src.Type),
		)
	}

	body := make([]UI, 0, src.Len())
	for i := 0; i < src.Len(); i++ {
		body = append(body, FilterUIElems(f(i))...)
	}

	r.body = body
	return r
}

func (r rangeLoop) Map(f func(string) UI) RangeLoop {
	src := reflect.ValueOf(r.source)
	if src.Kind() != reflect.Map {
		panic(errors.New("range loop source is not a map").
			Tag("src-type", src.Type),
		)
	}

	if keyType := src.Type().Key(); keyType.Kind() != reflect.String {
		panic(errors.New("range loop source keys are not strings").
			Tag("src-type", src.Type).
			Tag("key-type", keyType),
		)
	}

	body := make([]UI, 0, src.Len())
	keys := make([]string, 0, src.Len())

	for _, k := range src.MapKeys() {
		keys = append(keys, k.String())
	}
	sort.Strings(keys)

	for _, k := range keys {
		body = append(body, FilterUIElems(f(k))...)
	}

	r.body = body
	return r
}

func (r rangeLoop) Kind() Kind {
	return Selector
}

func (r rangeLoop) JSValue() Value {
	return nil
}

func (r rangeLoop) Mounted() bool {
	return false
}

func (r rangeLoop) name() string {
	return "range"
}

func (r rangeLoop) self() UI {
	return r
}

func (r rangeLoop) setSelf(UI) {
}

func (r rangeLoop) context() context.Context {
	return nil
}

func (r rangeLoop) attributes() map[string]string {
	return nil
}

func (r rangeLoop) eventHandlers() map[string]eventHandler {
	return nil
}

func (r rangeLoop) parent() UI {
	return nil
}

func (r rangeLoop) setParent(UI) {
}

func (r rangeLoop) children() []UI {
	return r.body
}

func (r rangeLoop) mount() error {
	return errors.New("range loop is not mountable").
		Tag("name", r.name()).
		Tag("kind", r.Kind())
}

func (r rangeLoop) dismount() {
}

func (r rangeLoop) update(UI) error {
	return errors.New("range loop cannot be updated").
		Tag("name", r.name()).
		Tag("kind", r.Kind())
}

func (r rangeLoop) onNav(*url.URL) {
}

// Condition represents a control structure that displays nodes depending on a
// given expression.
type Condition interface {
	UI

	// ElseIf sets the condition with the given nodes if previous expressions
	// were not met and given expression is true.
	ElseIf(expr bool, elems ...UI) Condition

	// Else sets the condition with the given UI elements if previous
	// expressions were not met.
	Else(elems ...UI) Condition
}

// If returns a condition that filters the given elements according to the given
// expression.
func If(expr bool, elems ...UI) Condition {
	if !expr {
		elems = nil
	}

	return condition{
		body:      FilterUIElems(elems...),
		satisfied: expr,
	}
}

type condition struct {
	body      []UI
	satisfied bool
}

func (c condition) ElseIf(expr bool, elems ...UI) Condition {
	if c.satisfied {
		return c
	}

	if expr {
		c.body = FilterUIElems(elems...)
		c.satisfied = expr
	}

	return c
}

func (c condition) Else(elems ...UI) Condition {
	return c.ElseIf(true, elems...)
}

func (c condition) Kind() Kind {
	return Selector
}

func (c condition) JSValue() Value {
	return nil
}

func (c condition) Mounted() bool {
	return false
}

func (c condition) name() string {
	return "if.else"
}

func (c condition) self() UI {
	return c
}

func (c condition) setSelf(UI) {
}

func (c condition) context() context.Context {
	return nil
}

func (c condition) attributes() map[string]string {
	return nil
}

func (c condition) eventHandlers() map[string]eventHandler {
	return nil
}

func (c condition) parent() UI {
	return nil
}

func (c condition) setParent(UI) {
}

func (c condition) children() []UI {
	return c.body
}

func (c condition) mount() error {
	panic("can't mout condition")
}

func (c condition) dismount() {
}

func (c condition) update(UI) error {
	panic("condition cannot be updated")
}
