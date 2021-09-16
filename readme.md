<p align="center">
    <a href="https://goreportcard.com/report/github.com/pyros2097/wapp"><img src="https://goreportcard.com/badge/github.com/pyros2097/wapp" alt="Go Report Card"></a>
	<a href="https://GitHub.com/pyros2097/wapp/releases/"><img src="https://img.shields.io/github/release/pyros2097/wapp.svg" alt="GitHub release"></a>
	<a href="https://pkg.go.dev/github.com/pyros2097/wapp"><img src="https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat" alt="pkg.go.dev docs"></a>
</p>

# wapp

**wapp** is a framework to build web apps in golang.
It uses a declarative syntax using funcs that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup.ZZZIt is highly opioninated and integrates uses tailwind css and alpinejs.

# Install

**wapp** requirements:

- [Go 1.16](https://golang.org/doc/go1.16)

```sh
go mod init
go get -u -v github.com/pyros2097/wapp
```

[Demo](https://github.com/pyros2097/wapp-example)

**Example**

```go
package main

import (
	. "github.com/pyros2097/wapp"
)

func Header() *Element {
	return Row(Css("w-full mb-20 font-bold text-xl text-gray-700 p-4"),
		Div(Css("text-blue-700"),
			A(Href("https://wapp.pyros2097.dev"), Text("wapp.pyros2097.dev")),
		),
		Div(Css("flex flex-row flex-1 justify-end items-end p-2"),
			Div(Css("border-b-2 border-white text-lg text-blue-700 mr-4"), Text("Examples: ")),
			Div(Css("link mr-4"), A(Href("/"), Text("Home"))),
			Div(Css("link mr-4"), A(Href("/clock"), Text("Clock"))),
			Div(Css("link mr-4"), A(Href("/about"), Text("About"))),
			Div(Css("link mr-4"), A(Href("/container"), Text("Container"))),
			Div(Css("link mr-4"), A(Href("/panic"), Text("Panic"))),
		),
	)
}

func Index(w http.ResponseWriter, r *http.Request) *Element {
	return Page(
		Col(
			Header(),
			H1(Text("Hello this is a h1")),
			H2(Text("Hello this is a h2")),
			H2(XData("{ message: 'I ❤️ Alpine' }"), XText("message"), Text("")),
			Col(Css("text-3xl text-gray-700"),
				Row(
					Row(Css("underline"),
						Text("Counter"),
					),
				),
				Row(
					Button(Css("btn m-20"),
						Text("-"),
					),
					Row(Css("m-20"),
						Text(strconv.Itoa(1)),
					),
					Button(Css("btn m-20"),
						Text("+"),
					),
				),
			),
		),
	)
}
```
