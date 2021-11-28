package components

import (
	"context"
	"strconv"

	. "github.com/pyros2097/gromer"
)

func Counter(c context.Context, start int) *Element {
	count, setCount := UseState(c, start)
	increment := func() {
		setCount(count().(int) + 1)
	}
	decrement := func() {
		setCount(count().(int) + 1)
	}
	return Col(Css("text-3xl text-gray-700"),
		Row(
			Row(Css("underline"),
				Text("Counter"),
			),
		),
		Row(
			Button2("-", "decrement"),
			Row(Css("m-20 text-5xl"), XText("count"),
				Text("count"),
			),
			Button2("+", "increment"),
		),
		M{
			"count":     strconv.Itoa(count().(int)),
			"increment": increment,
			"decrement": decrement,
		},
	)
}
