package components

import (
	"strconv"

	. "github.com/pyros2097/gromer"
)

// func CounterReducer(start int) Reducer {
// 	return Reducer{
// 		Name: "counter",
// 		State: State{
// 			"count": start,
// 		},
// 		Actions: Actions{
// 			"increment": func() string { return "this.state.count += 1;" },
// 			"decrement": func() string { return "this.state.count -= 1;" },
// 		},
// 	}
// }

func Counter(start int) *Element {
	// data := CounterReducer(start)
	return Col(Css("text-3xl text-gray-700"),
		Row(
			Row(Css("underline"),
				Text("Counter"),
			),
		),
		Row(
			Button2("-", "decrement"),
			Row(Css("m-20 text-5xl"), XText("state.count"),
				Text(strconv.Itoa(start)),
			),
			Button2("+", "increment"),
		),
	)
}
