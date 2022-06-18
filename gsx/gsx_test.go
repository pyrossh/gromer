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
	RegisterComponent(Todo, "todo")
	RegisterFunc(WebsiteName)
	h := Html(map[string]interface{}{
		"todos": []*TodoData{
			{ID: "b1a7359c-ebb4-11ec-8ea0-0242ac120002", Text: "My first todo", Completed: true},
		},
	})
	h["todo"] = &TodoData{ID: "b1a7359c-ebb4-11ec-8ea0-0242ac120002", Text: "My first todo", Completed: true}
	actual := h.Render(`
		<div>
			<div>
				123
				<ul id="todo-list" class="relative">
					<template x-for="todo in todos">
						<Todo key="todo">
							<div class="container">
								<h2>Title</h2>
								<h3>Sub title</h3>
							</div>
						</Todo>
					</template>
				</ul>
			</div>
			<div>
				Test
				<button>{WebsiteName}</button>
			</div>
		</div>
	`).String()
	expected := stripWhitespace(`
		<div>
			<div>
				123
				<ul id="todo-list" class="relative">
					<template x-for="todo in todos">
						<todo key="todo">
							<li id="todo-b1a7359c-ebb4-11ec-8ea0-0242ac120002" class="completed">
								<div class="view">
									<span>My first todo</span>
								</div>
								<div class="container">
									<h2>Title</h2>
									<h3>Sub title</h3>
								</div>
								<div class="count">
									<span>true</span>
								</div>
							</li>
						</todo>
					</template>
				</ul>
			</div>
			<div>
				Test
				<button>My Website</button>
			</div>
		</div>
	`)
	r.Equal(expected, actual)
}
