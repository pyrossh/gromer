package handlebars

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_blockParams(t *testing.T) {
	r := require.New(t)
	bp := NewBlockParams()
	r.Equal([]string{}, bp.current)
	r.Len(bp.stack, 0)

	bp.push([]string{"mark"})
	r.Equal([]string{"mark"}, bp.current)
	r.Len(bp.stack, 1)

	bp.push([]string{"bates"})
	r.Equal([]string{"bates"}, bp.current)
	r.Len(bp.stack, 2)
	r.Equal([][]string{
		[]string{"mark"},
		[]string{"bates"},
	}, bp.stack)

	b := bp.pop()
	r.Equal([]string{"bates"}, b)
	r.Equal([]string{"mark"}, bp.current)
	r.Len(bp.stack, 1)

	b = bp.pop()
	r.Equal([]string{"mark"}, b)
	r.Len(bp.stack, 0)
	r.Equal([]string{}, bp.current)
}

func newBlockParams() {
	panic("unimplemented")
}

func Test_Eval_Map_Call_Key(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	data := map[string]string{
		"a": "A",
		"b": "B",
	}
	ctx.Set("letters", data)
	s, _, err := Html(`
		{{letters.a}}|{{letters.b}}
	`).Props(
		"a", "A",
		"b", "B",
	).Render()
	r.NoError(err)
	r.Equal("A|B", strings.TrimSpace(string(s)))
}

func Test_Eval_Calls_on_Pointers(t *testing.T) {
	r := require.New(t)
	type user struct {
		Name string
	}
	u := &user{Name: "Mark"}
	ctx := NewContext()
	ctx.Set("user", u)

	s, err := Render("{{user.Name}}", ctx)
	r.NoError(err)
	r.Equal("Mark", s)
}
