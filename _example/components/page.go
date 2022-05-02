package components

import (
	"html/template"

	"github.com/pyros2097/gromer"
	"github.com/pyros2097/gromer/_example/assets"
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
					<meta name="author" content="pyrossh" />
					<meta name="keywords" content="pyros.sh, pyrossh, gromer"  />
					<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover" />
					<link rel="icon" href="{{ iconUrl }}" />
					<link rel="stylesheet" href="{{ todoCssUrl }}" />
					<link rel="stylesheet" href="{{ stylesCssUrl }}" />
					<script src="{{ htmxJsUrl }}"></script>
					<script src="{{ htmxJsonUrl }}"></script>
					<script src="{{ alpineJsUrl }}" defer=""></script>
			</head>
			<body>
			{{ props.Children }}
			</body>
		</html>
	`).Props(
		"iconUrl", gromer.GetAssetUrl(assets.FS, "images/icon.png"),
		"todoCssUrl", gromer.GetAssetUrl(assets.FS, "css/todo.css"),
		"stylesCssUrl", gromer.GetStylesUrl(),
		"htmxJsUrl", gromer.GetAssetUrl(assets.FS, "js/htmx@1.7.0.js"),
		"htmxJsonUrl", gromer.GetAssetUrl(assets.FS, "js/htmx.json-enc.js"),
		"alpineJsUrl", gromer.GetAssetUrl(assets.FS, "js/alpinejs@3.9.6.js"),
	)
}
