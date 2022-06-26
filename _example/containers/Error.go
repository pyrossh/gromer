package containers

import (
	. "github.com/pyros2097/gromer/gsx"
)

func Error(c Context, err error) *Node {
	c.Set("err", err.Error())
	return c.Render(`
		<span class="error">
			<strong>Failed to load: {err}</strong>
		</span>
	`)
}

var _ = Css(`
	.error {
		color: red;
		font-size: 18px;
		font-weight: 500;
		padding: 24px;
		border: red 1px solid;
	}
`)
