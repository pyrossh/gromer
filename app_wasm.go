package app

import (
	"net/url"

	"github.com/pyros2097/wapp/js"
)

var (
	body       *elem
	content    UI
	rootPrefix string
)

func Run(render RenderFunc) {
	defer func() {
		err := recover()
		// show alert
		panic(err)
	}()

	initBody()
	initContent()
	if err := body.replaceChildAt(0, render); err != nil {
		panic("replacing content failed")
	}
	content = render

	for {
		select {
		case f := <-uiChan:
			f()
		}
	}
}

func Reload() {
	dispatch(func() {
		js.Window.Location().Reload()
	})
}

func initBody() {
	body = &elem{
		jsvalue: js.Window.Get("document").Get("body"),
		tag:     "body",
	}
	body.setSelf(body)
}

func initContent() {
	content := &elem{
		jsvalue: body.JSValue().Get("firstElementChild"),
		tag:     "div",
	}

	content.setSelf(content)
	content.setParent(body)
	body.body = append(body.body, content)
}

func onPopState(this Value, args []Value) interface{} {
	dispatch(func() {
		// navigate(Window().URL(), false)
	})
	return nil
}

func isExternalNavigation(u *url.URL) bool {
	return u.Host != "" && u.Host != Window().URL().Host
}

func isFragmentNavigation(u *url.URL) bool {
	return u.Fragment != ""
}
