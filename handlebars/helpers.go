package handlebars

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// GlobalHelpers contains all of the default helpers for handlebars.
// These will be available to all templates. You should add
// any custom global helpers to this list.
var GlobalHelpers = HelperMap{
	moot:    &sync.Mutex{},
	helpers: map[string]interface{}{},
}

func init() {
	GlobalHelpers.Add("if", ifHelper)
	GlobalHelpers.Add("unless", unlessHelper)
	GlobalHelpers.Add("each", eachHelper)
	GlobalHelpers.Add("eq", equalHelper)
	GlobalHelpers.Add("equal", equalHelper)
	GlobalHelpers.Add("neq", notEqualHelper)
	GlobalHelpers.Add("notequal", notEqualHelper)
	GlobalHelpers.Add("json", toJSONHelper)
	GlobalHelpers.Add("js_escape", template.JSEscapeString)
	GlobalHelpers.Add("html_escape", template.HTMLEscapeString)
	GlobalHelpers.Add("upcase", strings.ToUpper)
	GlobalHelpers.Add("downcase", strings.ToLower)
	GlobalHelpers.Add("len", lenHelper)
	GlobalHelpers.Add("debug", debugHelper)
	GlobalHelpers.Add("inspect", inspectHelper)
}

// HelperContext is an optional context that can be passed
// as the last argument to helper functions.
type HelperContext struct {
	Context     *Context
	Args        []interface{}
	evalVisitor *evalVisitor
}

// Block executes the block of template associated with
// the helper, think the block inside of an "if" or "each"
// statement.
func (h HelperContext) Block() (string, error) {
	return h.BlockWith(h.Context)
}

// BlockWith executes the block of template associated with
// the helper, think the block inside of an "if" or "each"
// statement. It takes a new context with which to evaluate
// the block.
func (h HelperContext) BlockWith(ctx *Context) (string, error) {
	nev := newEvalVisitor(h.evalVisitor.template, ctx)
	nev.blockParams = h.evalVisitor.blockParams
	dd := nev.VisitProgram(h.evalVisitor.curBlock.Program)
	switch tp := dd.(type) {
	case string:
		return tp, nil
	case error:
		return "", errors.WithStack(tp)
	case nil:
		return "", nil
	default:
		return "", errors.WithStack(errors.Errorf("unknown return value %T %+v", dd, dd))
	}
}

// ElseBlock executes the "inverse" block of template associated with
// the helper, think the "else" block of an "if" or "each"
// statement.
func (h HelperContext) ElseBlock() (string, error) {
	return h.ElseBlockWith(h.Context)
}

// ElseBlockWith executes the "inverse" block of template associated with
// the helper, think the "else" block of an "if" or "each"
// statement. It takes a new context with which to evaluate
// the block.
func (h HelperContext) ElseBlockWith(ctx *Context) (string, error) {
	if h.evalVisitor.curBlock.Inverse == nil {
		return "", nil
	}
	nev := newEvalVisitor(h.evalVisitor.template, ctx)
	nev.blockParams = h.evalVisitor.blockParams
	dd := nev.VisitProgram(h.evalVisitor.curBlock.Inverse)
	switch tp := dd.(type) {
	case string:
		return tp, nil
	case error:
		return "", errors.WithStack(tp)
	case nil:
		return "", nil
	default:
		return "", errors.WithStack(errors.Errorf("unknown return value %T %+v", dd, dd))
	}
}

// Get is a convenience method that calls the underlying Context.
func (h HelperContext) Get(key string) interface{} {
	return h.Context.Get(key)
}

// toJSONHelper converts an interface into a string.
func toJSONHelper(v interface{}) (template.HTML, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return template.HTML(b), nil
}

func lenHelper(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	return strconv.Itoa(rv.Len())
}

// Debug by verbosely printing out using 'pre' tags.
func debugHelper(v interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("<pre>%+v</pre>", v))
}

func inspectHelper(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}

func eachHelper(collection interface{}, help HelperContext) (template.HTML, error) {
	out := bytes.Buffer{}
	val := reflect.ValueOf(collection)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct || val.Len() == 0 {
		s, err := help.ElseBlock()
		return template.HTML(s), err
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			ctx := help.Context.New()
			ctx.Set("@first", i == 0)
			ctx.Set("@last", i == val.Len()-1)
			ctx.Set("@index", i)
			ctx.Set("@value", v)
			s, err := help.BlockWith(ctx)
			if err != nil {
				return "", errors.WithStack(err)
			}
			out.WriteString(s)
		}
	case reflect.Map:
		keys := val.MapKeys()
		for i := 0; i < len(keys); i++ {
			key := keys[i].Interface()
			v := val.MapIndex(keys[i]).Interface()
			ctx := help.Context.New()
			ctx.Set("@first", i == 0)
			ctx.Set("@last", i == len(keys)-1)
			ctx.Set("@key", key)
			ctx.Set("@value", v)
			s, err := help.BlockWith(ctx)
			if err != nil {
				return "", errors.WithStack(err)
			}
			out.WriteString(s)
		}
	}
	return template.HTML(out.String()), nil
}

func equalHelper(a, b interface{}, help HelperContext) (template.HTML, error) {
	if a == b {
		s, err := help.Block()
		if err != nil {
			return "", err
		}
		return template.HTML(s), nil
	}
	s, err := help.ElseBlock()
	if err != nil {
		return "", err
	}
	return template.HTML(s), nil
}

func notEqualHelper(a, b interface{}, help HelperContext) (template.HTML, error) {
	if a != b {
		s, err := help.Block()
		if err != nil {
			return "", err
		}
		return template.HTML(s), nil
	}
	s, err := help.ElseBlock()
	if err != nil {
		return "", err
	}
	return template.HTML(s), nil
}

func ifHelper(conditional interface{}, help HelperContext) (template.HTML, error) {
	if IsTrue(conditional) {
		s, err := help.Block()
		return template.HTML(s), err
	}
	s, err := help.ElseBlock()
	return template.HTML(s), err
}

// IsTrue returns true if obj is a truthy value.
func IsTrue(obj interface{}) bool {
	thruth, ok := isTrueValue(reflect.ValueOf(obj))
	if !ok {
		return false
	}
	return thruth
}

// isTrueValue reports whether the value is 'true', in the sense of not the zero of its type,
// and whether the value has a meaningful truth value
//
// NOTE: borrowed from https://github.com/golang/go/tree/master/src/text/template/exec.go
func isTrueValue(val reflect.Value) (truth, ok bool) {
	if !val.IsValid() {
		// Something like var x interface{}, never set. It's a form of nil.
		return false, true
	}
	switch val.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		truth = val.Len() > 0
	case reflect.Bool:
		truth = val.Bool()
	case reflect.Complex64, reflect.Complex128:
		truth = val.Complex() != 0
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Interface:
		truth = !val.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		truth = val.Int() != 0
	case reflect.Float32, reflect.Float64:
		truth = val.Float() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		truth = val.Uint() != 0
	case reflect.Struct:
		truth = true // Struct values are always true.
	default:
		return
	}
	return truth, true
}

func unlessHelper(conditional bool, help HelperContext) (template.HTML, error) {
	return ifHelper(!conditional, help)
}

// HelperMap holds onto helpers and validates they are properly formed.
type HelperMap struct {
	moot    *sync.Mutex
	helpers map[string]interface{}
}

// Add a new helper to the map. New Helpers will be validated to ensure they
// meet the requirements for a helper:
/*
	func(...) (string) {}
	func(...) (string, error) {}
	func(...) (template.HTML) {}
	func(...) (template.HTML, error) {}
*/
func (h *HelperMap) Add(key string, helper interface{}) error {
	h.moot.Lock()
	defer h.moot.Unlock()
	err := h.validateHelper(key, helper)
	if err != nil {
		return errors.WithStack(err)
	}
	h.helpers[key] = helper
	return nil
}

func (h *HelperMap) validateHelper(key string, helper interface{}) error {
	ht := reflect.ValueOf(helper).Type()

	if ht.NumOut() < 1 {
		return errors.WithStack(errors.Errorf("%s must return at least one value ([string|template.HTML], [error])", key))
	}
	so := ht.Out(0).Kind().String()
	if ht.NumOut() > 1 {
		et := ht.Out(1)
		ev := reflect.ValueOf(et)
		ek := fmt.Sprintf("%s", ev.Interface())
		if (so != "string" && so != "template.HTML") || (ek != "error") {
			return errors.WithStack(errors.Errorf("%s must return ([string|template.HTML], [error]), not (%s, %s)", key, so, et.Kind()))
		}
	} else {
		if so != "string" && so != "template.HTML" {
			return errors.WithStack(errors.Errorf("%s must return ([string|template.HTML], [error]), not (%s)", key, so))
		}
	}
	return nil
}
