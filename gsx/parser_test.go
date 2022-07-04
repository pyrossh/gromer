package gsx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	r := require.New(t)
	actual := RenderString(parse("test", `
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
	actual := RenderString(parse("test", `
		<Todo />
		<TodoCount />
	`))
	expected := strings.TrimLeft(`
<Todo />
<TodoCount />
`, "\n")
	r.Equal(expected, actual)
}

func TestForLoop(t *testing.T) {
	r := require.New(t)
	actual := RenderString(parse("test", `
		<ul>
		for k, v := range todos {
			return (
				<li>
					"data"
				</li>
				<div>
					<span>
					{name}
					</span>
				</div>
			)
		}
		</ul>
	`))
	expected := strings.TrimLeft(`
<ul>
  <>
  </>
</ul>
`, "\n")
	r.Equal(expected, actual)
}
