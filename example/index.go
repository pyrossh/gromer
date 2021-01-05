package main

import (
	"strconv"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/components"
)

func Index(c *RenderContext) UI {
	count, setCount := c.UseInt(0)
	inc := func() { setCount(count() + 1) }
	dec := func() { setCount(count() - 1) }
	return Col(
		Header(c),
		Col(Css("text-3xl text-gray-700"),
			Row(
				Row(Css("underline"),
					Text("Counter"),
				),
			),
			Row(
				Button(Css("btn m-20"), OnClick(dec),
					Text("-"),
				),
				Row(Css("m-20"),
					Text(strconv.Itoa(count())),
				),
				Button(Css("btn m-20"), OnClick(inc),
					Text("+"),
				),
			),
		),
	)
}
