<p align="center">
    <a href="https://goreportcard.com/report/github.com/pyros2097/wapp"><img src="https://goreportcard.com/badge/github.com/pyros2097/wapp" alt="Go Report Card"></a>
	<a href="https://GitHub.com/pyros2097/wapp/releases/"><img src="https://img.shields.io/github/release/pyros2097/wapp.svg" alt="GitHub release"></a>
	<a href="https://pkg.go.dev/github.com/pyros2097/wapp/v7/pkg/app"><img src="https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat" alt="pkg.go.dev docs"></a>
</p>

**wapp** is a package to build [isomorphic web apps](https://developers.google.com/web/progressive-web-apps/) with [Go](https://golang.org) and [WebAssembly](https://webassembly.org).

It uses a [declarative syntax](#declarative-syntax) that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup. The syntax is inspired by react and its awesome hooks and functional component features.

This originally started out as of fork of this awesome golang PWA framework [go-app](https://github.com/pyros2097/wapp). All credits goes to Maxence Charriere for majority of the work.

## Install

**wapp** requirements:

- [Go 1.15](https://golang.org/doc/go1.15)

```sh
go mod init
go get -u -v github.com/pyros2097/wapp
```

## Declarative syntax

**go-app** uses a declarative syntax so you can write component-based UI elements just by using the Go programming language.

```go
package pages

import (
	"github.com/pyros2097/wapp"
)

func Index(c *app.RenderContext) app.UI {
	count, setCount := c.UseInt(0)
	onclick := func(ctx app.Context, e app.Event) {
		setCount(count() + 1)
	}

	return app.Div().
		Body(
			app.Div().
				Style("cursor", "pointer").
				OnClick(onclick).
				Body(
					app.Text(count()),
				),
			app.Text("id: "+id.(string)),
		)
}
```

The directory should look like as the following:

```sh
.                     
├── pages             
    └── index.go      
    └── about.go      
└── static            
    └── favicon.png   
    └── logo.png   
```
