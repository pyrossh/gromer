package handlebars

import (
	"sync"

	"github.com/aymerick/raymond/ast"
	"github.com/aymerick/raymond/parser"
	"github.com/pkg/errors"
)

// Template represents an input and helpers to be used
// to evaluate and render the input.
type Template struct {
	Input   string
	program *ast.Program
}

// NewTemplate from the input string.
func NewTemplate(input string) (*Template, error) {
	t := &Template{
		Input: input,
	}
	err := t.Parse()
	if err != nil {
		return t, errors.WithStack(err)
	}
	return t, nil
}

// Parse the template this can be called many times
// as a successful result is cached and is used on subsequent
// uses.
func (t *Template) Parse() error {
	if t.program != nil {
		return nil
	}
	program, err := parser.Parse(t.Input)
	if err != nil {
		return errors.WithStack(err)
	}
	t.program = program
	return nil
}

// Exec the template using the content and return the results
func (t *Template) Exec(ctx *Context) (string, error) {
	err := t.Parse()
	if err != nil {
		return "", errors.WithStack(err)
	}
	v := newEvalVisitor(t, ctx)
	r := t.program.Accept(v)
	switch rp := r.(type) {
	case string:
		return rp, nil
	case error:
		return "", rp
	case nil:
		return "", nil
	default:
		return "", errors.WithStack(errors.Errorf("unsupport eval return format %T: %+v", r, r))
	}
}

var cache = map[string]*Template{}
var moot = &sync.Mutex{}

// Parse an input string and return a Template.
func Parse(input string) (*Template, error) {
	moot.Lock()
	defer moot.Unlock()
	if t, ok := cache[input]; ok {
		return t, nil
	}
	t, err := NewTemplate(input)

	if err == nil {
		cache[input] = t
	}

	if err != nil {
		return t, errors.WithStack(err)
	}

	return t, nil
}

// Render a string using the given the context.
func Render(input string, ctx *Context) (string, error) {
	t, err := Parse(input)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return t.Exec(ctx)
}
