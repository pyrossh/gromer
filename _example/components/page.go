package components

import (
	. "github.com/pyros2097/gromer/gsx"
)

var _ = Css(`
	hr {
		margin: 20px 0;
		border: 0;
		border-top: 1px dashed #c5c5c5;
		border-bottom: 1px dashed #f7f7f7;
	}

	html,
	body {
		margin: 0;
		padding: 0;
	}

	button {
		margin: 0;
		padding: 0;
		border: 0;
		background: none;
		font-size: 100%;
		vertical-align: baseline;
		font-family: inherit;
		font-weight: inherit;
		color: inherit;
		-webkit-appearance: none;
		appearance: none;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	body {
		font: 14px 'Helvetica Neue', Helvetica, Arial, sans-serif;
		line-height: 1.4em;
		background: #f5f5f5;
		color: #4d4d4d;
		min-width: 230px;
		max-width: 550px;
		margin: 0 auto;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		font-weight: 300;
	}

	:focus {
		outline: 0;
	}

	.hidden {
		display: none;
	}

	.new-todo,
	.edit {
		position: relative;
		margin: 0;
		width: 100%;
		font-size: 24px;
		font-family: inherit;
		font-weight: inherit;
		line-height: 1.4em;
		border: 0;
		color: inherit;
		padding: 6px;
		border: 1px solid #999;
		box-shadow: inset 0 -1px 5px 0 rgba(0, 0, 0, 0.2);
		box-sizing: border-box;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	.new-todo {
		padding: 16px 16px 16px 60px;
		border: none;
		background: rgba(0, 0, 0, 0.003);
		box-shadow: inset 0 -2px 1px rgba(0, 0, 0, 0.03);
	}

	.main {
		position: relative;
		z-index: 2;
		border-top: 1px solid #e6e6e6;
	}

	.toggle-all {
		text-align: center;
		border: none; /* Mobile Safari */
		opacity: 0;
		position: absolute;
	}

	.toggle-all + label {
		width: 60px;
		height: 34px;
		font-size: 0;
		position: absolute;
		top: -52px;
		left: -13px;
		-webkit-transform: rotate(90deg);
		transform: rotate(90deg);
	}

	.toggle-all + label:before {
		content: '❯';
		font-size: 22px;
		color: #e6e6e6;
		padding: 10px 27px 10px 27px;
	}

	.toggle-all:checked + label:before {
		color: #737373;
	}

	.footer {
		color: #777;
		padding: 10px 15px;
		height: 20px;
		text-align: center;
		border-top: 1px solid #e6e6e6;
	}

	.footer:before {
		content: '';
		position: absolute;
		right: 0;
		bottom: 0;
		left: 0;
		height: 50px;
		overflow: hidden;
		box-shadow: 0 1px 1px rgba(0, 0, 0, 0.2), 0 8px 0 -3px #f6f6f6, 0 9px 1px -3px rgba(0, 0, 0, 0.2), 0 16px 0 -6px #f6f6f6, 0 17px 2px -6px rgba(0, 0, 0, 0.2);
	}

	.filters {
		margin: 0;
		padding: 0;
		list-style: none;
		position: absolute;
		right: 0;
		left: 0;
	}

	.filters li {
		display: inline;
	}

	.filters li a {
		color: inherit;
		margin: 3px;
		padding: 3px 7px;
		text-decoration: none;
		border: 1px solid transparent;
		border-radius: 3px;
	}

	.filters li a:hover {
		border-color: rgba(175, 47, 47, 0.1);
	}

	.filters li a.selected {
		border-color: rgba(175, 47, 47, 0.2);
	}

	.clear-completed,
	html .clear-completed:active {
		float: right;
		position: relative;
		line-height: 20px;
		text-decoration: none;
		cursor: pointer;
	}

	.clear-completed:hover {
		text-decoration: underline;
	}

	.info {
		margin: 65px auto 0;
		color: #bfbfbf;
		font-size: 10px;
		text-shadow: 0 1px 0 rgba(255, 255, 255, 0.5);
		text-align: center;
	}

	.info p {
		line-height: 1;
	}

	.info a {
		color: inherit;
		text-decoration: none;
		font-weight: 400;
	}

	.info a:hover {
		text-decoration: underline;
	}

	@media screen and (-webkit-min-device-pixel-ratio: 0) {
		.toggle-all {
			background: none;
		}
	}

	@media (max-width: 430px) {
		.footer {
			height: 50px;
		}

		.filters {
			bottom: 10px;
		}
	}
`)

func Page(h Html, title string) string {
	return h.Render(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
					<meta charset="UTF-8" />
					<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
					<meta content="utf-8" http-equiv="encoding" />
					<title>{title}</title>
					<meta name="description" content="{title}" />
					<meta name="author" content="pyrossh" />
					<meta name="keywords" content="pyros.sh, pyrossh, gromer"  />
					<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover" />
					<link rel="icon" href="{GetAssetUrl "images/icon.png"}" />
					<link rel="stylesheet" href="{GetStylesUrl}" />
					<script src="{GetAlpineJsUrl}"></script>
					<script src="{GetHtmxJsUrl}" defer=""></script>
			</head>
			<body>
			{children}
			</body>
		</html>
	`)
}
