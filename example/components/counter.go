package components

import (
	"strconv"

	. "github.com/pyros2097/wapp"
	. "github.com/pyros2097/wapp/example/context"
)

// data := map[string]interface{}{
// 	"count": 1,
// 	"names": []string{"123", "123"},
// 	"increment": func() string {
// 		return "this.count += 1;"
// 	},
// 	"decrement": func() string {
// 		return "this.count -= 1;"
// 	},
// }

func Counter(ctx ReqContext, start int) *Element {
	UseData(ctx, "counter", M{
		"count":     start,
		"increment": func() string { return "this.count += 1;" },
		"decrement": func() string { return "this.count -= 1;" },
	})
	return Col(Css("text-3xl text-gray-700"),
		Row(
			Row(Css("underline"),
				Text("Counter"),
			),
		),
		Row(XData("counter"),
			Button(Css("btn m-20"), OnClick("decrement"),
				Text("-"),
			),
			Row(Css("m-20 text-8xl"), XText("count"),
				Text(strconv.Itoa(start)),
			),
			Button(Css("btn m-20"), OnClick("increment"),
				Text("+"),
			),
		),
	)
}

// func HTML(s string, m map[string]interface{}) string {
// 	return s
// }

// func Counter2(start int) string {
// 	return HTML(`
// 		<div class="flex flex-col justify-center items-center text-3xl text-gray-700">
// 			<div class="flex flex-row justify-center items-center">
// 				<div class="flex flex-row justify-center items-center underline">
// 				Counter
// 				</div>
// 			</div>
// 			<div class="flex flex-row justify-center items-center" x-data="counter">
// 				<button @click="decrement" class="btn m-20">-</button>
// 				<div class="flex flex-row justify-center items-center m-20">{{ count }}</div>
// 				<button @click="increment" class="btn m-20">+</button>
// 			</div>
// 			<div>
// 			{{#if true }}
// 				render this
// 			{{/if}}
// 			</div>
// 			<div>
// 			{{#each names}}
// 				<li>{{ @index }} - {{ @value }}</li>
// 			{{/each}}
// 			</div>
// 		</div>
// 	`, data)
// }

// if becomes
// <template x-if="open">
// </template>

// each becomes
// <ul x-data="{ names: ['Red', 'Orange', 'Yellow'] }">
//     <template x-for="name in names">
//         <li x-text="names"></li>
//     </template>
// </ul>
