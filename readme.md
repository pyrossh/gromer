# gromer

[![Version](https://badge.fury.io/gh/pyros2097%2Fgromer.svg)](https://github.com/pyros2097/gromer)

**gromer** is a framework and cli to build web apps in golang.
It uses a declarative syntax using funcs that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup.

It also generates http handlers for your routes which follow a particular folder structure. Similar to other frameworks like nextjs, sveltekit.

# Install

```sh
go get -u -v github.com/pyros2097/gromer/cmd/gromer
```

You can install this plugin https://marketplace.visualstudio.com/items?itemName=pyros2097.vscode-go-inline-html for syntax highlighting html templates in golang.

# Using

You need to follow this directory structure similar to nextjs for the api route handlers to be generated

Take a look at the example for now,

[Example](https://github.com/pyros2097/gromer/tree/master/example)
