package components

import (
	"github.com/pyros2097/gromer/_example/services/todos"
	. "github.com/pyros2097/gromer/gsx"
)

var TodoStyles = M{
	"container": "border-t-2 border-gray-100 text-2xl",
	"row":       "flex flex-row group",
	"button-1":  "ml-4 text-gray-400",
	"label":     "flex-1 min-w-0 flex items-center break-all ml-2 p-2 text-gray-800",
	"striked":   "text-gray-500 line-through",
	"button-2":  "mr-4 text-red-700 opacity-0 hover:opacity-100",
	"unchecked": "text-gray-200",
}

func Todo(c *Context, todo *todos.Todo) []*Tag {
	checked := "/icons/unchecked.svg?fill=gray-400"
	if todo.Completed {
		checked = "/icons/checked.svg?fill=green-500"
	}
	c.Set("checked", checked)
	return c.Render(`
		<div id="todo-{todo.ID}" class="Todo">
			<div class="row">
				<form hx-post="/" hx-target="#todo-{todo.ID}" hx-swap="outerHTML">
					<input type="hidden" name="intent" value="complete" />
					<input type="hidden" name="id" value={todo.ID} />
					<button class="button-1">	
						<img src={checked} width="24" height="24" />
					</button>
				</form>
				<label class={ "label": true, "striked": todo.Completed }>
					{todo.Text}
				</label>
				<form hx-post="/" hx-target="#todo-{todo.ID}" hx-swap="delete">
					<input type="hidden" name="intent" value="delete" />
					<input type="hidden" name="id" value={todo.ID} />
					<button class="button-2">
						<img src="/icons/close.svg?fill=red-500" width="24" height="24" />
					</button>
				</form>
			</div>
		</div>	
	`)
}
