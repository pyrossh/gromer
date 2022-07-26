package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/pyros2097/gromer"
	_ "github.com/pyros2097/gromer/_example/routes"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) (*require.Assertions, playwright.Page, func()) {
	r := require.New(t)
	server := httptest.NewServer(gromer.GetRouter())
	pw, err := playwright.Run()
	r.NoError(err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	r.NoError(err)
	page, err := browser.NewPage()
	r.NoError(err)
	_, err = page.Goto(server.URL)
	r.NoError(err)
	return r, page, func() {
		server.Close()
		err = browser.Close()
		r.NoError(err)
		err = pw.Stop()
		r.NoError(err)
	}
}

func checkAttr(r *require.Assertions, el playwright.ElementHandle, exp string) {
	src, err := el.GetAttribute("src")
	r.NoError(err)
	r.Equal(exp, src)
}

func getEl(el playwright.ElementHandle, s string) playwright.ElementHandle {
	v, err := el.QuerySelector(s)
	if err != nil {
		panic(err)
	}
	return v
}

func TestTodoPage(t *testing.T) {
	r, page, close := setup(t)
	defer close()
	el, err := page.QuerySelector(".title")
	r.NoError(err)
	title, err := el.TextContent()
	r.NoError(err)
	r.Contains(title, "todos")
}

func TestAddTodo(t *testing.T) {
	r, page, close := setup(t)
	defer close()
	page.Type("#text", "First Todo")
	page.Keyboard().Down("Enter")
	els, err := page.QuerySelectorAll(".Todo")
	r.NoError(err)
	r.Len(els, 1)
	todo := els[0]
	title, err := todo.TextContent()
	r.NoError(err)
	r.Contains(title, "First Todo")
	img, err := todo.QuerySelector("img")
	r.NoError(err)
	checkAttr(r, img, "/icons/unchecked.svg?fill=gray-400")
	getEl(todo, ".button-1").Click()
	els, err = page.QuerySelectorAll(".Todo")
	r.NoError(err)
	checkAttr(r, getEl(els[0], "img"), "/icons/checked.svg?fill=green-500")
}
