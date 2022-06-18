package gsx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TodoData struct {
	ID        string
	Text      string
	Completed bool
}

func Todo(h Html, todo *TodoData) Node {
	return h.Render(`
		<li id="todo-{todo.ID}" class="{ completed: todo.Completed }">
			<div class="view">
				<span>{todo.Text}</span>
			</div>
			{children}
			<div class="count">
				<span>{todo.Completed}</span>
			</div>
		</li>
	`)
}

func WebsiteName() string {
	return "My Website"
}

func TestHtml(t *testing.T) {
	r := require.New(t)
	RegisterComponent("todo", Todo, "todo")
	RegisterFunc(WebsiteName)
	h := Html(map[string]interface{}{
		"todos": []*TodoData{
			{ID: "b1a7359c-ebb4-11ec-8ea0-0242ac120002", Text: "My first todo", Completed: true},
		},
	})
	h["todo"] = &TodoData{ID: "b1a7359c-ebb4-11ec-8ea0-0242ac120002", Text: "My first todo", Completed: true}
	// <template x-for="todo in todos">
	actual := h.Render(`
		<div>
			<div>
				123
				<Todo key="todo">
					<div class="container">
						<h2>Title</h2>
						<h3>Sub title</h3>
					</div>
				</Todo>
			</div>
			<div>
				Test
				<button>click</button>
			</div>
		</div>
	`).String()
	expected := "<div><div>123<todo key=\"todo\"><li id=\"todo-b1a7359c-ebb4-11ec-8ea0-0242ac120002\" class=\"{ completed: todo.Completed }\"><div class=\"view\"><span>{todo.Text}</span></div><div class=\"container\"><h2>Title</h2><h3>Sub title</h3></div><div class=\"count\"><span>{todo.Completed}</span></div></li></todo></div><div>Test<button>click</button></div></div>"
	// 	actual := Html(ctx).Render(`
	// 		<ul id="todo-list" class="relative">
	// 			<For key="todos" itemKey="todo">
	// 				<Todo key="todo">
	// 					<div>"Todo123"</div>
	// 				</Todo>
	// 			</For>
	// 			<span>{WebsiteName}</span>
	// 		</ul>
	// 	`)
	// 	expected := `<ul id="todo-list" class="relative">
	//   <For key="todos" itemKey="todo">
	//     <Todo key="todo">
	//       <li id="b1a7359c-ebb4-11ec-8ea0-0242ac120002" class="completed">
	//               <div>
	//         Todo123
	//       </div>
	//         <div class="view">
	//           <span>
	//             My first todo
	//           </span>
	//         </div>
	//       </li>
	//     </Todo>
	//   </For>
	//   <span>
	//     My Website
	//   </span>
	// </ul>`
	r.Equal(expected, actual)
}
