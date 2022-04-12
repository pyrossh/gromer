package handlebars

import (
	"fmt"
	"html/template"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CustomGlobalHelper(t *testing.T) {
	r := require.New(t)
	err := GlobalHelpers.Add("say", func(name string) (string, error) {
		return fmt.Sprintf("say: %s", name), nil
	})
	r.NoError(err)

	input := `{{say "mark"}}`
	ctx := NewContext()
	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("say: mark", s)
}

func Test_CustomGlobalBlockHelper(t *testing.T) {
	r := require.New(t)
	GlobalHelpers.Add("say", func(name string, help HelperContext) (template.HTML, error) {
		ctx := help.Context
		ctx.Set("name", strings.ToUpper(name))
		s, err := help.BlockWith(ctx)
		return template.HTML(s), err
	})

	input := `
	{{#say "mark"}}
		<h1>{{name}}</h1>
	{{/say}}
	`
	ctx := NewContext()
	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "<h1>MARK</h1>")
}

func Test_Helper_Hash_Options(t *testing.T) {
	r := require.New(t)
	GlobalHelpers.Add("say", func(help HelperContext) string {
		return help.Context.Get("name").(string)
	})

	input := `{{say name="mark"}}`
	ctx := NewContext()
	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("mark", s)
}

func Test_Helper_Hash_Options_Many(t *testing.T) {
	r := require.New(t)
	GlobalHelpers.Add("say", func(help HelperContext) string {
		return help.Context.Get("first").(string) + help.Context.Get("last").(string)
	})

	input := `{{say first=first_name last=last_name}}`
	ctx := NewContext()
	ctx.Set("first_name", "Mark")
	ctx.Set("last_name", "Bates")
	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("MarkBates", s)
}

func Test_Helper_Santize_Output(t *testing.T) {
	r := require.New(t)

	GlobalHelpers.Add("safe", func(help HelperContext) template.HTML {
		return template.HTML("<p>safe</p>")
	})
	GlobalHelpers.Add("unsafe", func(help HelperContext) string {
		return "<b>unsafe</b>"
	})

	input := `{{safe}}|{{unsafe}}`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Equal("<p>safe</p>|&lt;b&gt;unsafe&lt;/b&gt;", s)
}

func Test_JSON_Helper(t *testing.T) {
	r := require.New(t)

	input := `{{json names}}`
	ctx := NewContext()
	ctx.Set("names", []string{"mark", "bates"})
	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal(`["mark","bates"]`, s)
}

func Test_If_Helper(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	input := `{{#if true}}hi{{/if}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("hi", s)
}

func Test_If_Helper_false(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	input := `{{#if false}}hi{{/if}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("", s)
}

func Test_If_Helper_NoArgs(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	input := `{{#if }}hi{{/if}}`

	_, err := Render(input, ctx)
	r.Error(err)
}

func Test_If_Helper_Else(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	input := `
	{{#if false}}
		hi
	{{ else }}
		bye
	{{/if}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "bye")
}

func Test_Unless_Helper(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	input := `{{#unless false}}hi{{/unless}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("hi", s)
}

func Test_EqualHelper_True(t *testing.T) {
	r := require.New(t)
	input := `
	{{#eq 1 1}}
		it was true
	{{else}}
		it was false
	{{/eq}}
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "it was true")
}

func Test_EqualHelper_False(t *testing.T) {
	r := require.New(t)
	input := `
	{{#eq 1 2}}
		it was true
	{{else}}
		it was false
	{{/eq}}
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "it was false")
}

func Test_EqualHelper_DifferentTypes(t *testing.T) {
	r := require.New(t)
	input := `
	{{#eq 1 "1"}}
		it was true
	{{else}}
		it was false
	{{/eq}}
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "it was false")
}

func Test_NotEqualHelper_True(t *testing.T) {
	r := require.New(t)
	input := `
	{{#neq 1 1}}
		it was true
	{{else}}
		it was false
	{{/neq}}
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "it was false")
}

func Test_NotEqualHelper_False(t *testing.T) {
	r := require.New(t)
	input := `
	{{#neq 1 2}}
		it was true
	{{else}}
		it was false
	{{/neq}}
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "it was true")
}

func Test_NotEqualHelper_DifferentTypes(t *testing.T) {
	r := require.New(t)
	input := `
	{{#neq 1 "1"}}
		it was true
	{{else}}
		it was false
	{{/neq}}
	`
	s, err := Render(input, NewContext())
	r.NoError(err)
	r.Contains(s, "it was true")
}

func Test_Each_Helper_NoArgs(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	input := `{{#each }}{{@value}}{{/each}}`

	_, err := Render(input, ctx)
	r.Error(err)
}

func Test_Each_Helper(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	ctx.Set("names", []string{"mark", "bates"})
	input := `{{#each names }}<p>{{@value}}</p>{{/each}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("<p>mark</p><p>bates</p>", s)
}

func Test_Each_Helper_Index(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	ctx.Set("names", []string{"mark", "bates"})
	input := `{{#each names }}<p>{{@index}}</p>{{/each}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("<p>0</p><p>1</p>", s)
}

func Test_Each_Helper_As(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	ctx.Set("names", []string{"mark", "bates"})
	input := `{{#each names as |ind name| }}<p>{{ind}}-{{name}}</p>{{/each}}`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Equal("<p>0-mark</p><p>1-bates</p>", s)
}

func Test_Each_Helper_As_Nested(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	users := []struct {
		Name     string
		Initials []string
	}{
		{Name: "Mark", Initials: []string{"M", "F", "B"}},
		{Name: "Rachel", Initials: []string{"R", "A", "B"}},
	}
	ctx.Set("users", users)
	input := `
{{#each users as |user|}}
	<h1>{{user.Name}}</h1>
	{{#each user.Initials as |i|}}
		{{user.Name}}: {{i}}
	{{/each}}
{{/each}}
	`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "<h1>Mark</h1>")
	r.Contains(s, "Mark: M")
	r.Contains(s, "Mark: F")
	r.Contains(s, "Mark: B")
	r.Contains(s, "<h1>Rachel</h1>")
	r.Contains(s, "Rachel: R")
	r.Contains(s, "Rachel: A")
	r.Contains(s, "Rachel: B")
}

func Test_Each_Helper_SlicePtr(t *testing.T) {
	r := require.New(t)
	type user struct {
		Name string
	}
	type users []user

	us := &users{
		{Name: "Mark"},
		{Name: "Rachel"},
	}

	ctx := NewContext()
	ctx.Set("users", us)

	input := `
	{{#each users as |user|}}
		{{user.Name}}
	{{/each}}
	`
	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "Mark")
	r.Contains(s, "Rachel")
}

func Test_Each_Helper_Map(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	data := map[string]string{
		"a": "A",
		"b": "B",
	}
	ctx.Set("letters", data)
	input := `
	{{#each letters}}
		{{@key}}:{{@value}}
	{{/each}}
	`

	s, err := Render(input, ctx)
	r.NoError(err)
	for k, v := range data {
		r.Contains(s, fmt.Sprintf("%s:%s", k, v))
	}
}

func Test_Each_Helper_Map_As(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	data := map[string]string{
		"a": "A",
		"b": "B",
	}
	ctx.Set("letters", data)
	input := `
	{{#each letters as |k v|}}
		{{k}}:{{v}}
	{{/each}}
	`

	s, err := Render(input, ctx)
	r.NoError(err)
	for k, v := range data {
		r.Contains(s, fmt.Sprintf("%s:%s", k, v))
	}
}

func Test_Each_Helper_Else(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	data := map[string]string{}
	ctx.Set("letters", data)
	input := `
	{{#each letters as |k v|}}
		{{k}}:{{v}}
	{{else}}
		no letters
	{{/each}}
	`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "no letters")
}

func Test_Each_Helper_Else_Collection(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	data := map[string][]string{}
	ctx.Set("collection", data)

	input := `
	{{#each collection.emptykey as |k v|}}
		{{k}}:{{v}}
	{{else}}
		no letters
	{{/each}}
	`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "no letters")
}

func Test_Each_Helper_Else_CollectionMap(t *testing.T) {
	r := require.New(t)
	ctx := NewContext()
	data := map[string]map[string]string{
		"emptykey": map[string]string{},
	}

	ctx.Set("collection", data)

	input := `
	{{#each collection.emptykey.something as |k v|}}
		{{k}}:{{v}}
	{{else}}
		no letters
	{{/each}}
	`

	s, err := Render(input, ctx)
	r.NoError(err)
	r.Contains(s, "no letters")
}

func Test_HelperMap_Add(t *testing.T) {
	r := require.New(t)
	err := GlobalHelpers.Add("foo", func(help HelperContext) (string, error) {
		return "", nil
	})
	r.NoError(err)
}

func Test_HelperMap_Add_Invalid_NoReturn(t *testing.T) {
	r := require.New(t)
	err := GlobalHelpers.Add("foo", func(help HelperContext) {})
	r.Error(err)
	r.Contains(err.Error(), "must return at least one")
}

func Test_HelperMap_Add_Invalid_ReturnTypes(t *testing.T) {
	r := require.New(t)
	err := GlobalHelpers.Add("foo", func(help HelperContext) (string, string) {
		return "", ""
	})
	r.Error(err)
	r.Contains(err.Error(), "foo must return ([string|template.HTML], [error]), not (string, string)")

	err = GlobalHelpers.Add("foo", func(help HelperContext) int { return 1 })
	r.Error(err)
	r.Contains(err.Error(), "foo must return ([string|template.HTML], [error]), not (int)")
}
