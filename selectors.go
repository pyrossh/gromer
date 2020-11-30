package app

import (
	"reflect"
	"sort"
)

// Range returns a range loop that iterates within the given source. Source must
// be a slice, an array or a map with strings as keys.
func Range(src interface{}) RangeLoop {
	return RangeLoop{source: src}
}

// RangeLoop represents a control structure that iterates within a slice, an
// array or a map.
type RangeLoop struct {
	body   []UI
	source interface{}
}

// Slice sets the loop content by repeating the given function for the
// number of elements in the source.
//
// It panics if the range source is not a slice or an array.
func (r RangeLoop) Slice(f func(int) UI) RangeLoop {
	src := reflect.ValueOf(r.source)
	if src.Kind() != reflect.Slice && src.Kind() != reflect.Array {
		panic("range loop source is not a slice or array: " + src.Type().String())
	}

	body := make([]UI, 0, src.Len())
	for i := 0; i < src.Len(); i++ {
		body = append(body, FilterUIElems(f(i))...)
	}

	r.body = body
	return r
}

// Map sets the loop content by repeating the given function for the number
// of elements in the source. Elements are ordered by keys.
//
// It panics if the range source is not a map or if map keys are not strings.
func (r RangeLoop) Map(f func(string) UI) RangeLoop {
	src := reflect.ValueOf(r.source)
	if src.Kind() != reflect.Map {
		panic("range loop source is not a map: " + src.Type().String())
	}

	if keyType := src.Type().Key(); keyType.Kind() != reflect.String {
		panic("range loop source keys are not strings: " + src.Type().String() + keyType.String())
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

// If returns a condition that filters the given elements according to the given
// expression.
func If(expr bool, elems ...interface{}) Condition {
	if !expr {
		elems = nil
	}

	return Condition{
		body:      FilterUIElems(elems...),
		satisfied: expr,
	}
}

// Condition represents a control structure that displays nodes depending on a
// given expression.
type Condition struct {
	body      []UI
	satisfied bool
}

// ElseIf sets the condition with the given nodes if previous expressions
// were not met and given expression is true.
func (c Condition) ElseIf(expr bool, elems ...interface{}) Condition {
	if c.satisfied {
		return c
	}

	if expr {
		c.body = FilterUIElems(elems...)
		c.satisfied = expr
	}

	return c
}

// Else sets the condition with the given UI elements if previous
// expressions were not met.
func (c Condition) Else(elems ...interface{}) Condition {
	return c.ElseIf(true, elems...)
}
