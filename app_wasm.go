package app

import (
	"context"
	"net/url"
)

var (
	body       *elem
	content    UI
	rootPrefix string
)

func run(render RenderFunc) {
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

func initBody() {
	ctx, cancel := context.WithCancel(context.Background())
	body = &elem{
		ctx:       ctx,
		ctxCancel: cancel,
		jsvalue:   Window().Get("document").Get("body"),
		tag:       "body",
	}
	body.setSelf(body)
}

func initContent() {
	ctx, cancel := context.WithCancel(context.Background())

	content := &elem{
		ctx:       ctx,
		ctxCancel: cancel,
		jsvalue:   body.JSValue().Get("firstElementChild"),
		tag:       "div",
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
