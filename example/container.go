package main

import (
	"strconv"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/atoms"
)

func AtomCounter(c *RenderContext, no string) UI {
	count := c.UseAtom(CountAtom)
	return Col(
		Row(
			Row(Css("text-2xl"),
				Text("Counter - "+no),
			),
		),
		Row(
			Row(Css("text-2xl m-20 cursor-pointer select-none"), OnClick(DecCount),
				Text("-"),
			),
			Row(Css("text-2xl m-20"),
				Text(strconv.Itoa(count.(int))),
			),
			Row(Css("text-2xl m-20 cursor-pointer select-none"), OnClick(IncCount),
				Text("+"),
			),
		),
	)
}

func Container(c *RenderContext) UI {
	return Col(
		AtomCounter(c, "1"),
		AtomCounter(c, "2"),
		AtomCounter(c, "3"),
	)
}
