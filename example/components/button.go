package components

import (
	. "github.com/pyros2097/gromer"
)

func Button2(title, clickHandler string) *Element {
	return Button(Css("bg-gray-300 border-b-2 border-white hover:bg-gray-200 focus:outline-none rounded text-gray-700"), OnClick(clickHandler),
		Div(Css("flex flex-row flex-1 justify-center items-center p-4"),
			Text(title),
		),
	)
}
