package components

import (
	"context"
	"strconv"

	. "github.com/pyros2097/gromer"
)

type S map[string]interface{}

type CC func(nodes ...interface{}) *Element

func Styled2(s S) CC {
	return func(nodes ...interface{}) *Element {
		return Div(nodes...)
	}
}

var Container = Styled2(S{
	"border-left":  "2px",
	"border-right": "2px",
})

func Counter(c context.Context, start int) *Element {
	count, setCount := UseState(c, start)
	increment := func() {
		setCount(count().(int) + 1)
	}
	decrement := func() {
		setCount(count().(int) + 1)
	}
	return Container("123", Css("text-3xl text-gray-700"),
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
