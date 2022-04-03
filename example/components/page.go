package components

import (
	. "github.com/pyros2097/gromer"
)

func Page() string {
	return Component(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
					<meta charset="UTF-8" />
					<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
					<meta content="utf-8" http-equiv="encoding" />
					<title>{{ title }}</title>
					<meta name="description" content="{{ title }}" />
					<meta name="author" content="pyros.sh" />
					<meta content="pyros.sh, gromer" name="keywords" />
					<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover" />
					<link rel="icon" href="/assets/icon.png" />
					<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css" />
					<script src="https://unpkg.com/htmx.org@1.7.0"></script>
					<script src="/assets/alpine.js" defer=""></script>
			</head>
			<body>
			{{ children }}
			</body>
		</html>
	`)
}
