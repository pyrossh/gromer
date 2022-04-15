package handlebars

import (
	"fmt"

	"github.com/aymerick/raymond/ast"
	"github.com/aymerick/raymond/parser"
	"github.com/pkg/errors"
)

var stylesCss = CssContent("")

type CssContent string
type HtmlContent string

// Template represents an input and helpers to be used
// to evaluate and render the input.
type Template struct {
	Input   string
	Context *Context
	program *ast.Program
}

// NewTemplate from the input string.
func NewTemplate(input string) (*Template, error) {
	t := &Template{
		Input:   input,
		Context: NewContext(),
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
func (t *Template) RenderWithStatus(status int) (HtmlContent, int, error) {
	err := t.Parse()
	if err != nil {
		return HtmlContent("Server Erorr"), 500, errors.WithStack(err)
	}
	v := newEvalVisitor(t, t.Context)
	r := t.program.Accept(v)
	switch rp := r.(type) {
	case string:
		return HtmlContent(rp), status, nil
	case error:
		return HtmlContent("Server Erorr"), 500, rp
	case nil:
		return HtmlContent(""), 200, nil
	default:
		return HtmlContent("Server Erorr"), 500, errors.WithStack(errors.Errorf("unsupport eval return format %T: %+v", r, r))
	}
}

func (t *Template) Render() (HtmlContent, int, error) {
	return t.RenderWithStatus(200)
}

func (t *Template) Prop(key string, v any) *Template {
	t.Context.Set(key, v)
	return t
}

func (t *Template) Props(args ...any) *Template {
	for i := 0; i < len(args); i += 2 {
		key := fmt.Sprintf("%s", args[i])
		t.Context.Set(key, args[i+1])
	}
	return t
}

func Html(tpl string) *Template {
	return &Template{
		Input:   tpl,
		Context: NewContext(),
	}
}

func HtmlErr(status int, err error) (HtmlContent, int, error) {
	return HtmlContent("ErrorPage/AccessDeniedPage/NotFoundPage based on status code"), status, err
}

func Css(v string) CssContent {
	stylesCss += CssContent(v)
	return CssContent(v)
}

func GetStyles() CssContent {
	return stylesCss
}
