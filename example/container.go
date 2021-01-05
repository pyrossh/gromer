package main

import (
	"strconv"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/atoms"
	. "github.com/pyros2097/wapp/example/components"
)

func AtomCounter(c *RenderContext, no string) UI {
	count := c.UseAtom(CountAtom)
	return Col(Css("text-3xl text-gray-700"),
		Row(
			Row(Css("underline"),
				Text("Counter - "+no),
			),
		),
		Row(
			Button(Css("btn m-20"), OnClick(DecCount),
				Text("-"),
			),
			Row(Css("m-20"),
				Text(strconv.Itoa(count.(int))),
			),
			Button(Css("btn m-20"), OnClick(IncCount),
				Text("+"),
			),
		),
	)
}

func Container(c *RenderContext) UI {
	return Col(
		Header(c),
		AtomCounter(c, "1"),
		AtomCounter(c, "2"),
		AtomCounter(c, "3"),
	)
}
