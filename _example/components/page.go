package components

import (
	"html/template"

	. "github.com/pyros2097/gromer/handlebars"
)

type PageProps struct {
	Title    string        `json:"title"`
	Children template.HTML `json:"children"`
}

func Page(props PageProps) *Template {
	return Html(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
					<meta charset="UTF-8" />
					<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
					<meta content="utf-8" http-equiv="encoding" />
					<title>{{ props.Title }}</title>
					<meta name="description" content="{{ props.Title }}" />
					<meta name="author" content="pyros.sh" />
					<meta content="pyros.sh, gromer" name="keywords" />
					<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover" />
					<link rel="icon" href="/assets/icon.png" />
					<link rel="stylesheet" href="/assets/css/bulma@0.9.3.css" />
					<link rel="stylesheet" href="/styles.css" />
					<script src="/assets/js/htmx@1.7.0.js"></script>
					<script src="/assets/js/alpinejs@3.9.6.js" defer=""></script>
			</head>
			<body>
			{{ props.Children }}
			</body>
		</html>
	`)
}
