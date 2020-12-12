<p align="center">
    <a href="https://goreportcard.com/report/github.com/pyros2097/wapp"><img src="https://goreportcard.com/badge/github.com/pyros2097/wapp" alt="Go Report Card"></a>
	<a href="https://GitHub.com/pyros2097/wapp/releases/"><img src="https://img.shields.io/github/release/pyros2097/wapp.svg" alt="GitHub release"></a>
	<a href="https://pkg.go.dev/github.com/pyros2097/wapp"><img src="https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat" alt="pkg.go.dev docs"></a>
</p>

**wapp** is a framework to build isomorphic web apps in golang.

It uses a declarative syntax using funcs that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup. The syntax is inspired by react and its awesome hooks and functional component features. It is highly opioninated and integrates very well with tailwind css for now.

This originally started out as of fork of this awesome go-app PWA framework. All credits goes to Maxence Charriere for majority of the work.

Inspired by:
* [go-app](https://github.com/maxence-charriere/go-app)
* [react](https://reactjs.org/docs/components-and-props.html)
* [reacct-hooks](https://reactjs.org/docs/hooks-intro.html)
* [jotai](https://github.com/pmndrs/jotai)
* [klyva](https://github.com/merisbahti/klyva)


## Install

**wapp** requirements:

- [Go 1.15](https://golang.org/doc/go1.15)

```sh
go mod init
go get -u -v github.com/pyros2097/wapp
```

## Example

**wapp** uses a declarative syntax so you can write component-based UI elements just by using the Go programming language. It follows the same ideas of react. It has functional components and hooks.

The example is located here,

[example](https://github.com/pyros2097/wapp/tree/master/example)

**Counter**

```go
package main

import (
	. "github.com/pyros2097/wapp"
)

func Counter(c *RenderContext) UI {
	count, setCount := c.UseInt(0)
	inc := func() { setCount(count() + 1) }
	dec := func() { setCount(count() - 1) }
	return Col(
		Row(
			Row(Css("text-6xl"),
				Text("Counter"),
			),
		),
		Row(
			Row(Css("text-6xl m-20 cursor-pointer select-none"), OnClick(dec),
				Text("-"),
			),
			Row(Css("text-6xl m-20"),
				Text(count()),
			),
			Row(Css("text-6xl m-20 cursor-pointer select-none"), OnClick(inc),
				Text("+"),
			),
		),
	)
}

func main() {
	Route("/", Counter)
	Run(Counter)
}
```

**Clock**

```go
package main

import (
	"time"

	. "github.com/pyros2097/wapp"
)

func Route(c *RenderContext, title string) UI {
	timeValue, setTime := c.UseState(time.Now())
	running, setRunning := c.UseState(false)
	startTimer := func() {
		setRunning(true)
		go func() {
			for running().(bool) {
				setTime(time.Now())
				time.Sleep(time.Second * 1)
			}
		}()
	}
	stopTimer := func() {
		setRunning(false)
	}
	c.UseEffect(func() func() {
		startTimer()
		return stopTimer
	})

	return Col(
		Row(
			Div(Css("text-6xl"),
				Text(title),
			),
		),
		Row(
			Div(Css("mt-10"),
				Text(timeValue().(time.Time).Format("15:04:05")),
			),
		),
		Row(
			Div(Css("text-6xl m-20 cursor-pointer select-none"), OnClick(startTimer),
				Text("Start"),
			),
			Div(Css("text-6xl m-20 cursor-pointer select-none"), OnClick(stopTimer),
				Text("Stop"),
			),
		),
	)
}

func main() {
	Route("/", Clock)
	Run(Clock)
}

```
