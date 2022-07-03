package gsx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	r := require.New(t)
	actual := renderString(parse("test", `
<ul id="todo-list" class="relative">
	<Todo todo={v}>
		<div>"Todo123"</div>
	</Todo>
	<img src="123" />
	<span>{WebsiteName}</span>
</ul>
<div>
	<p>		
	"Done"
	</p>
</div>
	`))
	println(actual)
	expected := strings.TrimLeft(`
<ul id=""todo-list"" class=""relative"">
  <Todo>
    <div>
      Todo123
    </div>
  </Todo>
  <img src=""123"" />
  <span>
    {WebsiteName}
  </span>
</ul>
<div>
  <p>
    Done
  </p>
</div>
`, "\n")
	r.Equal(expected, actual)
}

func TestSelfClose(t *testing.T) {
	r := require.New(t)
	actual := renderString(parse("test", `
		<Todo />
		<TodoCount />
	`))
	println(actual)
	expected := strings.TrimLeft(`
<Todo />
<TodoCount />
`, "\n")
	r.Equal(expected, actual)
}
