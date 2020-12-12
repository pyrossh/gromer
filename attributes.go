package app

import (
	"github.com/pyros2097/wapp/js"
)

type CssAttribute struct {
	classes string
}

func Css(d string) CssAttribute {
	return CssAttribute{classes: d}
}

type OnClickAttribute struct {
	cb func()
}

func OnClick(cb func()) OnClickAttribute {
	return OnClickAttribute{cb: cb}
}

type OnChangeAttribute struct {
	cb js.EventHandlerFunc
}

func OnChange(cb js.EventHandlerFunc) OnChangeAttribute {
	return OnChangeAttribute{cb: cb}
}

type OnInputAttribute struct {
	cb js.EventHandlerFunc
}

func OnInput(cb js.EventHandlerFunc) OnInputAttribute {
	return OnInputAttribute{cb: cb}
}

type HelmetTitle string
type HelmetDescription string
type HelmetAuthor string
type HelmetKeywords string

func mergeAttributes(parent *elem, uis ...interface{}) {
	elems := []UI{}
	for _, v := range uis {
		switch c := v.(type) {
		case CssAttribute:
			if vv, ok := parent.attrs["classes"]; ok {
				parent.setAttr("class", vv+" "+c.classes)
			} else {
				parent.setAttr("class", c.classes)
			}
		case OnClickAttribute:
			parent.setEventHandler("click", func(e js.Event) {
				c.cb()
			})
		case OnChangeAttribute:
			parent.setEventHandler("change", c.cb)
		case OnInputAttribute:
			parent.setEventHandler("input", c.cb)
		case HelmetTitle:
			helmet.Title = string(c)
		case HelmetDescription:
			helmet.Description = string(c)
		case HelmetAuthor:
			helmet.Author = string(c)
		case HelmetKeywords:
			helmet.Keywords = string(c)
		case UI:
			elems = append(elems, c)
		default:
			panic("unknown type in render")
		}
	}
	parent.setBody(elems)
}
