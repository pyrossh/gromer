package gromer

import (
	"bytes"
	"context"
	"strconv"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	. "github.com/franela/goblin"
)

func TestGetRouteParams(t *testing.T) {
	g := Goblin(t)
	g.Describe("GetRouteParams", func() {
		g.It("should return all the right params", func() {
			params := GetRouteParams("/api/todos/{id}/update/{action}")
			g.Assert(params).Equal([]string{"id", "action"})
		})
	})
}

func Row(uis ...interface{}) *Element {
	return NewElement("div", false, append([]interface{}{Css("flex flex-row justify-center items-center")}, uis...)...)
}

func Col(uis ...interface{}) *Element {
	return NewElement("div", false, append([]interface{}{Css("flex flex-col justify-center items-center")}, uis...)...)
}

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
			Button(OnClick("decrement"), Text("-")),
			Row(Css("m-20 text-5xl"), XText("count"),
				Text(strconv.Itoa(count().(int))),
			),
			Button(OnClick("increment"), Text("+")),
		),
		M{
			"count":     count,
			"increment": increment,
			"decrement": decrement,
		},
	)
	// return `
	// 	<div class="flex flex-col justify-center items-center text-3xl text-gray-700">
	// 		<div class="flex flex-row justify-center items-center">
	// 			<div class="flex flex-row justify-center items-center underline">
	// 				Counter
	// 			</div>
	// 		</div>
	// 		<div class="flex flex-row justify-center items-center">
	// 			<button class="btn m-20" @click="Increment">
	// 				-
	// 			</button>
	// 			<div class="flex flex-row justify-center items-center m-20 text-8xl">
	// 				{{ count }}
	// 			</div>
	// 			<button class="btn m-20" @click="Decrement">
	// 				+
	// 			</button>
	// 		</div>
	// 	</div>
	// `
}

func TestHtml(t *testing.T) {
	g := Goblin(t)
	g.Describe("Html", func() {
		g.It("should match snapshot", func() {
			ctx := WithState(context.Background())
			b := bytes.NewBuffer(nil)
			p := Html(
				Head(
					Title("123"),
					Meta("description", "123"),
					Meta("author", "123"),
					Meta("keywords", "123"),
					Meta("viewport", "width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover"),
					Link("icon", "/assets/icon.png"),
					Link("apple-touch-icon", "/assets/icon.png"),
					Link("stylesheet", "/assets/styles.css"),
					Script(Src("/assets/alpine.js"), Defer()),
					Meta("title", "title"),
				),
				Body(
					H1(Text("Hello this is a h1")),
					H2(Text("Hello this is a h2")),
					H3(XData("{ message: 'I ❤️ Alpine' }"), XText("message"), Text("")),
					Counter(ctx, 4),
				),
			)
			p.WriteHtml(b)
			c := cupaloy.New(cupaloy.SnapshotFileExtension(".html"))
			c.SnapshotT(t, b.String())
		})
	})
}
