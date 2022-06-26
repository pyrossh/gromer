package containers

import (
	. "github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

func TodoList(c Context, page int, filter string) *Node {
	index := Default(page, 1)
	todos, err := todos.GetAllTodo(c, todos.GetAllTodoParams{
		Filter: filter,
		Limit:  index,
	})
	if err != nil {
		return Error(c, err)
	}
	c.Set("todos", todos)
	return c.Render(`
		<ul id="todo-list" class="relative" x-for="todo in todos">
			<Todo />
		</ul>
	`)
}

var _ = Css(`
	.todo-list {
		margin: 0;
		padding: 0;
		list-style: none;
	}

	.todo-list li {
		position: relative;
		font-size: 24px;
		border-bottom: 1px solid #ededed;
	}

	.todo-list li:last-child {
		border-bottom: none;
	}

	.todo-list li.editing {
		border-bottom: none;
		padding: 0;
	}

	.todo-list li.editing .edit {
		display: block;
		width: 506px;
		padding: 12px 16px;
		margin: 0 0 0 43px;
	}

	.todo-list li.editing .view {
		display: none;
	}

	.todo-list li label {
		word-break: break-all;
		padding: 15px 15px 15px 60px;
		display: block;
		line-height: 1.2;
		transition: color 0.4s;
	}

	.todo-list li.completed label {
		color: #d9d9d9;
		text-decoration: line-through;
	}

	.todo-list li .destroy {
		display: none;
		position: absolute;
		top: 0;
		right: 10px;
		bottom: 0;
		width: 40px;
		height: 40px;
		margin: auto 0;
		font-size: 30px;
		color: #cc9a9a;
		margin-bottom: 11px;
		transition: color 0.2s ease-out;
	}

	.todo-list li .destroy:hover {
		color: #af5b5e;
	}

	.todo-list li .destroy:after {
		content: '×';
	}

	.todo-list li:hover .destroy {
		display: block;
	}

	.todo-list li .edit {
		display: none;
	}

	.todo-list li.editing:last-child {
		margin-bottom: -1px;
	}
	
	@media screen and (-webkit-min-device-pixel-ratio: 0) {
		.todo-list li .toggle {
			background: none;
		}

		.todo-list li .toggle {
			height: 40px;
		}
	}
`)
