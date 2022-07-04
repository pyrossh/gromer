package gsx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TodoData struct {
	ID        string
	Text      string
	Completed bool
}

func Todo(c *Context, todo *TodoData) []*Tag {
	return c.Render(`
		<li id="todo-{todo.ID}" class={"completed": todo.Completed }>
			<div class="upper">
				<span>{todo.Text}</span>
				<span>{todo.Text}</span>
			</div>
			{children}
			<div class="bottom">
				<span>{todo.Completed}</span>
				<span>{todo.Completed}</span>
			</div>
		</li>
	`)
}

func TodoList(c *Context, todos []*TodoData) []*Tag {
	return c.Render(`
		<ul id="todo-list" class="relative" x-for="todo in todos">
			<Todo />
		</ul>
	`)
}

func TodoCount(c *Context, count int) []*Tag {
	return c.Render(`
		<span id="todo-count" class="todo-count" hx-swap-oob="true">
			<strong>{count}</strong> "item left"
		</span>
	`)
}

func WebsiteName() string {
	return "My Website"
}

func trimLeft(s string) string {
	return strings.TrimLeft(s, "\n")
}

func TestComponent(t *testing.T) {
	r := require.New(t)
	RegisterComponent(Todo, nil, "todo")
	RegisterFunc(WebsiteName)
	h := Context{
		data: M{
			"todo": &TodoData{ID: "4", Text: "My fourth todo", Completed: true},
		},
	}
	nodes := h.Render(`
		<Todo>
			<span>{todo.Text}</span>
			<span>{todo.Completed}</span>
		</Todo>
		<Todo />
	`)
	actual := renderString(nodes)
	expected := trimLeft(`
<Todo>
  <li id="todo-4" class="completed">
    <div class="upper">
      <span>
        My fourth todo
      </span>
      <span>
        My fourth todo
      </span>
    </div>
    <span>
      My fourth todo
    </span>
    <span>
      true
    </span>

    <div class="bottom">
      <span>
        true
      </span>
      <span>
        true
      </span>
    </div>
  </li>
</Todo>
<Todo>
  <li id="todo-4" class="completed">
    <div class="upper">
      <span>
        My fourth todo
      </span>
      <span>
        My fourth todo
      </span>
    </div>

    <div class="bottom">
      <span>
        true
      </span>
      <span>
        true
      </span>
    </div>
  </li>
</Todo>
`)
	r.Equal(expected, actual)
}

func TestMultipleComponent(t *testing.T) {
	r := require.New(t)
	RegisterComponent(Todo, nil, "todo")
	RegisterComponent(TodoCount, nil, "count")
	h := Context{
		data: M{
			"todo":  &TodoData{ID: "4", Text: "My fourth todo", Completed: true},
			"count": 10,
		},
	}
	nodes := h.Render(`
			<Todo />
			<TodoCount />
	`)
	actual := renderString(nodes)
	expected := trimLeft(`
<Todo>
  <li id="todo-4" class="completed">
    <div class="upper">
      <span>
        My fourth todo
      </span>
      <span>
        My fourth todo
      </span>
    </div>

    <div class="bottom">
      <span>
        true
      </span>
      <span>
        true
      </span>
    </div>
  </li>
</Todo>
<TodoCount>
  <span id="todo-count" class="todo-count" hx-swap-oob="true">
    <strong>
      10
    </strong>
    item left
  </span>
</TodoCount>
`)
	r.Equal(expected, actual)
}

func TestFor(t *testing.T) {
	r := require.New(t)
	RegisterComponent(Todo, nil, "todo")
	RegisterFunc(WebsiteName)
	h := Context{
		data: map[string]interface{}{
			"todos": []*TodoData{
				{ID: "1", Text: "My first todo", Completed: true},
				{ID: "2", Text: "My second todo", Completed: false},
				{ID: "3", Text: "My third todo", Completed: false},
			},
		},
	}
	nodes := h.Render(`
		<ul class="relative">
			for i, v := range todos {
				return (
					<li>
						<span>{v.Text}</span>
						<span>{v.Completed}</span>
						<a>"link to" {v.ID}</a>
					</li>
				)
			}
		</ul>
		<ol>
			for i, v := range todos {
				return (
					<Todo todo={v}>
						<div class="todo-panel">
							<span>{v.Text}</span>
							<span>{v.Completed}</span>
						</div>
					</Todo>
				)
			}
		</ol>
	`)
	actual := renderString(nodes)
	expected := trimLeft(`
<ul class="relative">
  <li>
    <span>
      My first todo
    </span>
    <span>
      true
    </span>
    <a>
      link to
      1
    </a>
  </li>
  <li>
    <span>
      My second todo
    </span>
    <span>
      false
    </span>
    <a>
      link to
      2
    </a>
  </li>
  <li>
    <span>
      My third todo
    </span>
    <span>
      false
    </span>
    <a>
      link to
      3
    </a>
  </li>

</ul>
<ol>
  <Todo todo="v">
    <li id="todo-1" class="completed">
      <div class="upper">
        <span>
          My first todo
        </span>
        <span>
          My first todo
        </span>
      </div>
      <div class="todo-panel">
        <span>
          My first todo
        </span>
        <span>
          true
        </span>
      </div>

      <div class="bottom">
        <span>
          true
        </span>
        <span>
          true
        </span>
      </div>
    </li>
  </Todo>
  <Todo todo="v">
    <li id="todo-2">
      <div class="upper">
        <span>
          My second todo
        </span>
        <span>
          My second todo
        </span>
      </div>
      <div class="todo-panel">
        <span>
          My second todo
        </span>
        <span>
          false
        </span>
      </div>

      <div class="bottom">
        <span>
          false
        </span>
        <span>
          false
        </span>
      </div>
    </li>
  </Todo>
  <Todo todo="v">
    <li id="todo-3">
      <div class="upper">
        <span>
          My third todo
        </span>
        <span>
          My third todo
        </span>
      </div>
      <div class="todo-panel">
        <span>
          My third todo
        </span>
        <span>
          false
        </span>
      </div>

      <div class="bottom">
        <span>
          false
        </span>
        <span>
          false
        </span>
      </div>
    </li>
  </Todo>

</ol>
`)
	r.Equal(expected, actual)
}

// func TestForComponent(t *testing.T) {
// 	r := require.New(t)
// 	RegisterComponent(Todo, nil, "todo")
// 	RegisterComponent(TodoList, nil, "todos")
// 	RegisterFunc(WebsiteName)
// 	h := Context{
// 		data: map[string]interface{}{
// 			"todos": []*TodoData{
// 				{ID: "1", Text: "My first todo", Completed: true},
// 				{ID: "2", Text: "My second todo", Completed: false},
// 				{ID: "3", Text: "My third todo", Completed: false},
// 			},
// 		},
// 	}
// 	actual := h.Render(`
// 		<div>
// 			<TodoList />
// 		</div>
// 	`).String()
// 	expected := stripWhitespace(`
// 		<div>
// 			<ul id="todo-list" class="relative" x-for="todo in todos">
// 				<li id="todo-1" class="completed">
// 					<div class="view"><span>My first todo</span><span>My first todo</span></div>
// 					<div class="todo-panel"><span>My first todo</span><span>true</span></div>
// 					<div class="count"><span>true</span><span>true</span></div>
// 				</li>
// 				<li id="todo-2" class="">
// 					<div class="view"><span>My second todo</span><span>My second todo</span></div>
// 					<div class="todo-panel"><span>My second todo</span><span>false</span></div>
// 					<div class="count"><span>false</span><span>false</span></div>
// 				</li>
// 				<li id="todo-3" class="">
// 					<div class="view"><span>My third todo</span><span>My third todo</span></div>
// 					<div class="todo-panel"><span>My third todo</span><span>false</span></div>
// 					<div class="count"><span>false</span><span>false</span></div>
// 				</li>
// 			</ul>
// 		</div>
// 	`)
// 	r.Equal(expected, actual)
// }
