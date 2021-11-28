package components

import (
	. "github.com/pyros2097/gromer"
)

func Row(uis ...interface{}) *Element {
	return NewElement("div", false, append([]interface{}{Css("flex flex-row justify-center items-center")}, uis...)...)
}

func Col(uis ...interface{}) *Element {
	return NewElement("div", false, append([]interface{}{Css("flex flex-col justify-center items-center")}, uis...)...)
}
