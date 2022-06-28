package components

import (
	. "github.com/pyros2097/gromer/gsx"
)

// var CheckboxStyles = M{}

func Checkbox(c *Context, value bool) *Node {
	return c.Render(`
		<input class="checkbox" type="checkbox" checked="{value}" />
	`)
}
