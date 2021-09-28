package wapp

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func Row(uis ...interface{}) *Element {
	return NewElement("div", false, append([]interface{}{Css("flex flex-row justify-center items-center")}, uis...)...)
}

func Col(uis ...interface{}) *Element {
	return NewElement("div", false, append([]interface{}{Css("flex flex-col justify-center items-center")}, uis...)...)
}

// type Container func() *Element

// func CounterData(start int) Container {
// 	return CreateContainer(Reducer{
// 		Name: "counter",
// 		State: State{
// 			"count": start,
// 		},
// 		Actions: Actions{
// 			"increment": func() string { return "this.state.count += 1;" },
// 			"decrement": func() string { return "this.state.count -= 1;" },
// 		},
// 	})
// }

// type CounterState struct {
// 	Count int
// }

// type CounterActions struct {
// 	Increment func() string
// 	Decrement func() string
// }

// type CounterReducer struct {
// 	Reducer
// 	State   CounterState
// 	Actions CounterActions
// }

// func CounterData(start int) CounterReducer {
// 	return CounterReducer{
// 		Reducer: "counter",
// 		State: CounterState{
// 			Count: 1,
// 		},
// 		Actions: CounterActions{
// 			Increment: func() string { return "this.state.count += 1;" },
// 			Decrement: func() string { return "this.state.count -= 1;" },
// 		},
// 	}
// }

func CounterData(start int) Reducer {
	return Reducer{
		Name: "counter",
		State: State{
			"count": 1,
		},
		Actions: Actions{
			"increment": func() string { return "this.state.count += 1;" },
			"decrement": func() string { return "this.state.count -= 1;" },
		},
	}
}

func Counter(start int) *Element {
	data := CounterData(start)
	return Component(data, Col(Css("text-3xl text-gray-700"),
		Row(
			Row(Css("underline"),
				Text("Counter"),
			),
		),
		Row(XData("counter"),
			Button(Css("btn m-20"), OnClick("decrement"),
				Text("-"),
			),
			Row(Css("m-20 text-8xl"), XText("state.count"),
				Text(strconv.Itoa(start)),
			),
			Button(Css("btn m-20"), OnClick("increment"),
				Text("+"),
			),
		),
	))
}

func TestHtmlPage(t *testing.T) {
	b := bytes.NewBuffer(nil)
	p := Html(
		Head(
			Meta("title", "title"),
		),
		Body(
			H1(Text("Hello this is a h1")),
			H2(Text("Hello this is a h2")),
			H3(XData("{ message: 'I ❤️ Alpine' }"), XText("message"), Text("")),
			Counter(4),
		),
	)
	p.WriteHtml(b)
	cupaloy.SnapshotT(t, b.String())
}
