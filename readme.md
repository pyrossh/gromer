# gromer

[![Version](https://badge.fury.io/gh/pyros2097%2Fgromer.svg)](https://github.com/pyros2097/gromer)

**gromer** is a framework and cli to build web apps in golang.
It uses a declarative syntax using funcs that allows creating and dealing with HTML components and pages using templates.

It also generates http handlers for your routes which follow a particular folder structure. Similar to other frameworks like nextjs, sveltekit.
These handlers are also normal functions and can be imported in other packages directly ((inspired by [Encore](https://encore.dev/)).

# Install

```sh
go get -u -v github.com/pyros2097/gromer/cmd/gromer
```

# Using

You need to follow this directory structure similar to nextjs for the api route handlers to be generated

Take a look at the example for now,

[Example](https://github.com/pyros2097/gromer/tree/master/_example)

# Templating

Gromer uses a handlebars like templating language for components and pages. This is a modified version of this package [velvet](https://github.com/gobuffalo/velvet)
If you know handlebars, you basically know how to use it.

You can install this plugin [VSCode Go inline html plugin](https://marketplace.visualstudio.com/items?itemName=pyros2097.vscode-go-inline-html) for syntax highlighting the templates.

Let's assume you have a template (a string of some kind):

```handlebars
<!-- some input -->
<h1>{{ name }}</h1>
<ul>
  {{#each names}}
    <li>{{ @value }}</li>
  {{/each}}
</ul>
```

Given that string, you can render the template like such:

```html
<h1>Mark</h1>
<ul>
  <li>John</li>
  <li>Paul</li>
  <li>George</li>
  <li>Ringo</li>
</ul>
```

### If Statements

```handlebars
{{#if true }}
  render this
{{/if}}
```

#### Else Statements

```handlebars
{{#if false }}
  won't render this
{{ else }}
  render this
{{/if}}
```

#### Unless Statements

```handlebars
{{#unless true }}
  won't render this
{{/unless}}
```

### Each Statements

#### Arrays

When looping through `arrays` or `slices`, the block being looped through will be access to the "global" context, as well as have four new variables available within that block:

* `@first` [`bool`] - is this the first pass through the iteration?
* `@last` [`bool`] - is this the last pass through the iteration?
* `@index` [`int`] - the counter of where in the loop you are, starting with `0`.
* `@value` - the current element in the array or slice that is being iterated over.

```handlebars
<ul>
  {{#each names}}
    <li>{{ @index }} - {{ @value }}</li>
  {{/each}}
</ul>
```

By using "block parameters" you can change the "key" of the element being accessed from `@value` to a key of your choosing.

```handlebars
<ul>
  {{#each names as |name|}}
    <li>{{ name }}</li>
  {{/each}}
</ul>
```

To change both the key and the index name you can pass two "block parameters"; the first being the new name for the index and the second being the name for the element.

```handlebars
<ul>
  {{#each names as |index, name|}}
    <li>{{ index }} - {{ name }}</li>
  {{/each}}
</ul>
```

#### Maps

Looping through `maps` using the `each` helper is also supported, and follows very similar guidelines to looping through `arrays`.

* `@first` [`bool`] - is this the first pass through the iteration?
* `@last` [`bool`] - is this the last pass through the iteration?
* `@key` - the key of the pair being accessed.
* `@value` - the value of the pair being accessed.

```handlebars
<ul>
  {{#each users}}
    <li>{{ @key }} - {{ @value }}</li>
  {{/each}}
</ul>
```

By using "block parameters" you can change the "key" of the element being accessed from `@value` to a key of your choosing.

```handlebars
<ul>
  {{#each users as |user|}}
    <li>{{ @key }} - {{ user }}</li>
  {{/each}}
</ul>
```

To change both the key and the value name you can pass two "block parameters"; the first being the new name for the key and the second being the name for the value.

```handlebars
<ul>
  {{#each users as |key, user|}}
    <li>{{ key }} - {{ user }}</li>
  {{/each}}
</ul>
```
