<p align="center">
    <a href="https://goreportcard.com/report/github.com/pyros2097/gromer"><img src="https://goreportcard.com/badge/github.com/pyros2097/gromer" alt="Go Report Card"></a>
	<a href="https://GitHub.com/pyros2097/gromer/releases/"><img src="https://img.shields.io/github/release/pyros2097/gromer.svg" alt="GitHub release"></a>
	<a href="https://pkg.go.dev/github.com/pyros2097/gromer"><img src="https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat" alt="pkg.go.dev docs"></a>
</p>

# gromer

**gromer** is a framework and cli to build web apps in golang.
It uses a declarative syntax using funcs that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup.

It also generates http handlers for your routes which follow a particular folder structure. Similar to other frameworks like nextjs, sveltekit.

# Install

```sh
go get -u -v github.com/pyros2097/wapp/cmd/gromer
```

# Using

You need to follow this directory structure similar to nextjs for the api route handlers to be generated

Take a look at the example for now,

[Example](https://github.com/pyros2097/gromer/tree/master/example)
