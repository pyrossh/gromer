package gromer

import (
	"context"
	"encoding/json"
	"html/template"
	"strings"

	. "github.com/pyros2097/gromer/handlebars"
)

func ApiExplorer(ctx context.Context) (HtmlContent, int, error) {
	apiRoutes := []RouteDefinition{}
	for _, v := range RouteDefs {
		if strings.Contains(v.Path, "/api/") {
			apiRoutes = append(apiRoutes, v)
		}
	}
	apiData, err := json.Marshal(apiRoutes)
	if err != nil {
		return HtmlErr(400, err)
	}
	return Html(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
				<meta http-equiv="encoding" content="utf-8">
				<title> API Explorer </title>
				<meta name="description" content="API Explorer">
				<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover">
				<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/codemirror@5.63.1/lib/codemirror.css">
				<style>
					html,
					body {
						height: 100vh;
					}

					#left .CodeMirror {
						height: 400px;
					}

					#right .CodeMirror {
						height: calc(100vh - 60px);
					}

					.form-select {
						background-image: url("data:image/svg+xml,%3csvg
		xmlns='http://www.w3.org/2000/svg'viewBox='0 0 24 24'fill='%23a0aec0'%3e%3cpath d='M15.3 9.3a1 1 0 0 1 1.4 1.4l-4 4a1 1 0 0 1-1.4 0l-4-4a1 1 0 0 1 1.4-1.4l3.3 3.29 3.3-3.3z'/%3e%3c/svg%3e");
		-webkit-appearance: none;
								-moz-appearance: none;
								appearance: none;
								-webkit-print-color-adjust: exact;
								color-adjust: exact;
								background-repeat: no-repeat;
								background-color: #fff;
								border-color: #e2e8f0;
								border-width: 1px;
								border-radius: 0.25rem;
								padding-top: 0.5rem;
								padding-right: 2.5rem;
								padding-bottom: 0.5rem;
								padding-left: 0.75rem;
								font-size: 1rem;
								line-height: 1.5;
								background-position: right 0.5rem center;
								background-size: 1.5em 1.5em;
						}

						.form-select:focus {
							outline: none;
							box-shadow: 0 0 0 3px rgba(66, 153, 225, 0.5);
							border-color: #63b3ed;
						}

						table {
							width: 100%;
						}

						tr {
							width: 100%;
						}

						td, th {
							border-bottom: 1px solid rgb(204, 204, 204);
							border-left: 1px solid rgb(204, 204, 204);
							text-align: left;
						}

						textarea:focus, input:focus {
							outline: none;
							box-shadow: 0 0 0 3px rgba(66, 153, 225, 0.5);
							border-color: #63b3ed;
						}

						*:focus {
							outline: none;
							box-shadow: 0 0 0 3px rgba(66, 153, 225, 0.5);
							border-color: #63b3ed;
						}

						.spinner {
							animation: rotate 2s linear infinite;
							width: 24px;
							height: 24px;
						}

						.spinner .path {
							stroke: rgba(249, 250, 251, 1);
							stroke-linecap: round;
							animation: dash 1.5s ease-in-out infinite;
						}

						@keyframes rotate {
							100% {
								transform: rotate(360deg);
							}
						}

						@keyframes dash {
							0% {
								stroke-dasharray: 1, 150;
								stroke-dashoffset: 0;
							}

							50% {
								stroke-dasharray: 90, 150;
								stroke-dashoffset: -35;
							}

							100% {
								stroke-dasharray: 90, 150;
								stroke-dashoffset: -124;
							}
						}
				</style>
				<script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.63.1/codemirror.min.js"></script>
				<script src="https://cdn.jsdelivr.net/npm/codemirror@5.63.1/mode/javascript/javascript.js"></script>
				<style>
					*,
					::before,
					::after {
						box-sizing: border-box;
					}

					html {
						-moz-tab-size: 4;
						-o-tab-size: 4;
						tab-size: 4;
						line-height: 1.15;
						-webkit-text-size-adjust: 100%;
					}

					body {
						margin: 0;
						font-family: system-ui, -apple-system, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji';
					}

					hr {
						height: 0;
						color: inherit;
					}

					abbr[title] {
						-webkit-text-decoration: underline dotted;
						text-decoration: underline dotted;
					}

					b,
					strong {
						font-weight: bolder;
					}

					code,
					kbd,
					samp,
					pre {
						font-family: ui-monospace, SFMono-Regular, Consolas, 'Liberation Mono', Menlo, monospace;
						font-size: 1em;
					}

					small {
						font-size: 80%;
					}

					sub,
					sup {
						font-size: 75%;
						line-height: 0;
						position: relative;
						vertical-align: baseline;
					}

					sub {
						bottom: -0.25em;
					}

					sup {
						top: -0.5em;
					}

					table {
						text-indent: 0;
						border-color: inherit;
					}

					button,
					input,
					optgroup,
					select,
					textarea {
						font-family: inherit;
						font-size: 100%;
						line-height: 1.15;
						margin: 0;
					}

					button,
					select {
						text-transform: none;
					}

					button,
					[type='button'],
					[type='reset'],
					[type='submit'] {
						-webkit-appearance: button;
					}

					::-moz-focus-inner {
						border-style: none;
						padding: 0;
					}

					:-moz-focusring {
						outline: 1px dotted ButtonText;
						outline: auto;
					}

					:-moz-ui-invalid {
						box-shadow: none;
					}

					legend {
						padding: 0;
					}

					progress {
						vertical-align: baseline;
					}

					::-webkit-inner-spin-button,
					::-webkit-outer-spin-button {
						height: auto;
					}

					[type='search'] {
						-webkit-appearance: textfield;
						outline-offset: -2px;
					}

					::-webkit-search-decoration {
						-webkit-appearance: none;
					}

					::-webkit-file-upload-button {
						-webkit-appearance: button;
						font: inherit;
					}

					summary {
						display: list-item;
					}

					blockquote,
					dl,
					dd,
					h1,
					h2,
					h3,
					h4,
					h5,
					h6,
					hr,
					figure,
					p,
					pre {
						margin: 0;
					}

					button {
						background-color: transparent;
						background-image: none;
					}

					fieldset {
						margin: 0;
						padding: 0;
					}

					ol,
					ul {
						list-style: none;
						margin: 0;
						padding: 0;
					}

					html {
						font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
						line-height: 1.5;
					}

					body {
						font-family: inherit;
						line-height: inherit;
					}

					*,
					::before,
					::after {
						box-sizing: border-box;
						border-width: 0;
						border-style: solid;
						border-color: currentColor;
					}

					hr {
						border-top-width: 1px;
					}

					img {
						border-style: solid;
					}

					textarea {
						resize: vertical;
					}

					input::-moz-placeholder,
					textarea::-moz-placeholder {
						opacity: 1;
						color: #9ca3af;
					}

					input:-ms-input-placeholder,
					textarea:-ms-input-placeholder {
						opacity: 1;
						color: #9ca3af;
					}

					input::placeholder,
					textarea::placeholder {
						opacity: 1;
						color: #9ca3af;
					}

					button,
					[role="button"] {
						cursor: pointer;
					}

					table {
						border-collapse: collapse;
					}

					h1,
					h2,
					h3,
					h4,
					h5,
					h6 {
						font-size: inherit;
						font-weight: inherit;
					}

					a {
						color: inherit;
						text-decoration: inherit;
					}

					button,
					input,
					optgroup,
					select,
					textarea {
						padding: 0;
						line-height: inherit;
						color: inherit;
					}

					pre,
					code,
					kbd,
					samp {
						font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
					}

					img,
					svg,
					video,
					canvas,
					audio,
					iframe,
					embed,
					object {
						display: block;
						vertical-align: middle;
					}

					img,
					video {
						max-width: 100%;
						height: auto;
					}

					[hidden] {
						display: none;
					}

					*,
					::before,
					::after {
						--tw-border-opacity: 1;
						border-color: rgba(229, 231, 235, var(--tw-border-opacity));
					}

					.flex {
						display: flex;
					}

					.flex-col {
						flex-direction: column;
					}

					.w-full {
						width: 100%;
					}

					.p-2 {
						padding: 0.5rem;
					}

					.bg-gray-50 {
						background-color: rgba(249, 250, 251, 1);
					}

					.border-b {
						border-bottom-width: 1px;
					}

					.border-gray-200 {
						border-color: rgba(229, 231, 235, 1);
					}

					.items-center {
						align-items: center;
					}

					.justify-start {
						justify-content: flex-start;
					}

					.mr-4 {
						margin-right: 1rem;
					}

					.text-gray-700 {
						color: rgba(55, 65, 81, 1);
					}

					.text-2xl {
						font-size: 1.5rem;
						line-height: 2rem;
					}

					.font-bold {
						font-weight: 700;
					}

					.text-xl {
						font-size: 1.25rem;
						line-height: 1.75rem;
					}

					.block {
						display: block;
					}

					.ml-3 {
						margin-left: 0.75rem;
					}

					.mr-3 {
						margin-right: 0.75rem;
					}

					.bg-gray-200 {
						background-color: rgba(229, 231, 235, 1);
					}

					.border {
						border-width: 1px;
					}

					.border-gray-400 {
						border-color: rgba(156, 163, 175, 1);
					}

					.rounded-md {
						border-radius: 0.375rem;
					}

					.pt-2 {
						padding-top: 0.5rem;
					}

					.pb-2 {
						padding-bottom: 0.5rem;
					}

					.pl-6 {
						padding-left: 1.5rem;
					}

					.pr-6 {
						padding-right: 1.5rem;
					}

					.flex-row {
						flex-direction: row;
					}

					.pr-8 {
						padding-right: 2rem;
					}

					.border-r {
						border-right-width: 1px;
					}

					.border-gray-300 {
						border-color: rgba(209, 213, 219, 1);
					}

					.text-sm {
						font-size: 0.875rem;
						line-height: 1.25rem;
					}

					.uppercase {
						text-transform: uppercase
					}

					.pl-2 {
						padding-left: 0.5rem;
					}

					.p-1 {
						padding: 0.25rem;
					}

					.border-l {
						border-left-width: 1px;
					}

					.border-l-gray-200 {
						border-left-color: rgba(229, 231, 235, 1);
					}
				</style>
			</head>
			<body>
				<div class="flex flex-col">
					<div class="flex w-full p-2 bg-gray-50 border-b border-gray-200 items-center justify-start">
						<div class="flex mr-4 text-gray-700 text-2xl font-bold"> API Explorer </div>
						<div class="text-xl">
							<select id="api-select" class="form-select block">
								{{#each routes as |route|}}
									<option value="{{ @index }}">
										<div> {{ route.Method }} {{ route.Path }} </div>
									</option>
								{{/each}}
							</select>
						</div>
						<div class="flex ml-3 mr-3">
							<button id="run" class="bg-gray-200 border border-gray-400 hover:bg-gray-200 focus:outline-none rounded-md text-gray-700 text-md font-bold pt-2 pb-2 pl-6 pr-6"> RUN </button>
						</div>
					</div>
					<div class="flex">
						<div class="flex flex-row" style="width: 50%;;">
							<div class="pr-8 border-r border-gray-300" style="background: #f7f7f7;;"></div>
							<div class="w-full">
								<div class="text-gray-700 text-sm font-bold uppercase pl-2 pt-2 pb-2 bg-gray-50 border-b border-gray-200"> Headers </div>
								<table id="headersTable">
									<tr>
										<td>
											<input class="w-full p-1" value="Authorization">
										</td>
										<td>
											<input class="w-full p-1">
										</td>
									</tr>
								</table>
								<div class="text-gray-700 text-sm font-bold uppercase pl-2 pt-2 pb-2 bg-gray-50 border-b border-gray-200"> Path Params </div>
								<table id="pathParamsTable">
									<tr>
										<td class="text-gray-700" style="width: 50%;;">
											<div class="p-1"> 123 </div>
										</td>
										<td style="width: 50%;;">
											<input class="w-full p-1">
										</td>
									</tr>
								</table>
								<div class="text-gray-700 text-sm font-bold uppercase pl-2 pt-2 pb-2 bg-gray-50 border-b border-gray-200"> Query Params </div>
								<table id="queryParamsTable">
									<tr>
										<td class="text-gray-700" style="width: 50%;;">
											<div class="p-1"> 123 </div>
										</td>
										<td style="width: 50%;;">
											<input class="w-full p-1">
										</td>
									</tr>
								</table>
								<div class="text-gray-700 text-sm font-bold uppercase pl-2 pt-2 pb-2 bg-gray-50 border-b border-gray-200"> Body </div>
								<div id="left" class="border-b border-gray-200 text-md"></div>
							</div>
						</div>
						<div class="flex flex-row" style="width: 50%;;">
							<div id="right" class="w-full border-l border-l-gray-200 text-md"></div>
							<div class="pr-8 border-l border-gray-300" style="background: #f7f7f7;;"></div>
						</div>
					</div>
				</div>
				<script>
					window.apiDefs = {{ apiData }}
				</script>
				<script>
					window.codeLeft = CodeMirror(document.getElementById('left'), {
						value: '{}',
						mode: 'javascript'
					})
				</script>
				<script>
					window.codeRight = CodeMirror(document.getElementById('right'), {
						value: '',
						mode: 'javascript',
						lineNumbers: true,
						readOnly: true,
						lineWrapping: true
					})
				</script>
				<script>
					const getCurrentApiCall = () => {
						const index = document.getElementById("api-select").value;
						return window.apiDefs[index];
					}
					const updatePathParams = (apiCall) => {
						const table = document.getElementById("pathParamsTable");
						if (apiCall.pathParams.length === 0) {
							table.innerHTML = " <div style='background-color: rgb(245, 245, 245); padding: 0.25rem; text-align: center; color: gray;'>NONE</div>";
						} else {
							table.innerHTML = "";
						}
						for (const param of apiCall.pathParams.reverse()) {
							const row = table.insertRow(0);
							const cell1 = row.insertCell(0);
							const cell2 = row.insertCell(1);
							cell1.style = "width: 30%; border-left: 0px;";
							cell1.class = "text-gray-700";
							cell2.style = "width: 70%;";
							cell1.innerHTML = " <div class='p-1'> " + param + " </div>";
							cell2.innerHTML = " <input id='path-param-" + param + " class='w-full p-1'>";
						}
					}
					const updateParams = (apiCall) => {
						const table = document.getElementById("queryParamsTable");
						if (!apiCall.params) {
							table.innerHTML = " <div style='background-color: rgb(245, 245, 245); padding: 0.25rem; text-align: center; color: gray;'> NONE </div>";
					} else {
							table.innerHTML = "";
						}
						if (apiCall.method === "GET" || apiCall.method === "DELETE") {
							for (const key of Object.keys(apiCall.params)) {
								const row = table.insertRow(0);
								const cell1 = row.insertCell(0);
								const cell2 = row.insertCell(1);
								cell1.style = "width: 30%; border-left: 0px;";
								cell1.class = "text-gray-700";
								cell2.style = "width: 70%;";
								cell1.innerHTML = " <div class='p-1'> " + key + " </div>";
								cell2.innerHTML = " <input id='query-param-" + key + "' class='w-full p-1'>";
							}
						}
					}
					const updateBody = (apiCall) => {
						if (apiCall.method !== "GET" && apiCall.method !== "DELETE") {
							window.codeLeft.setValue(JSON.stringify(apiCall.params, 2, 2));
						} else {
							window.codeLeft.setValue("");
						}
					}
					const init = () => {
						updatePathParams(window.apiDefs[0]);
						updateParams(window.apiDefs[0]);
						updateBody(window.apiDefs[0]);
						const headersJson = localStorage.getItem("headers");
						if (headersJson) {
							const table = document.getElementById("headersTable");
							const headers = JSON.parse(headersJson);
							table.innerHTML = "";
							for (const key of Object.keys(headers)) {
								const value = headers[key];
								const row = table.insertRow(0);
								const cell1 = row.insertCell(0);
								const cell2 = row.insertCell(1);
								cell1.style = "width: 30%; border-left: 0px;";
								cell2.style = "width: 70%;";
								cell1.innerHTML = "<input value = '" + key + "' class='w-full p-1'>";
								cell2.innerHTML = "<input value = '" + value + "' class='w-full p-1'>";
							}
						}
					}
					window.onload = () => {
						init();
					}
					document.getElementById("api-select").onchange = () => {
						const apiCall = getCurrentApiCall();
						updatePathParams(apiCall);
						updateParams(apiCall);
						updateBody(apiCall);
					}
					const run = document.getElementById("run");
					run.onclick = async () => {
						run.innerHTML = "<svg class='spinner' viewBox='0 0 50 50'><circle class='path' cx='25' cy='25' r='20' fill='none' stroke-width='5'></circle></svg>";
						const table = document.getElementById("headersTable");
						const headers = {};
						for (const row of table.rows) {
							const key = row.cells[0].children[0].value;
							const value = row.cells[1].children[0].value;
							headers[key] = value;
						}
						const apiCall = getCurrentApiCall();
						let path = apiCall.path;
						const bodyParams = {};
						if (apiCall.method !== "GET" && apiCall.method != "DELETE") {
							bodyParams["body"] = window.codeLeft.getValue();
						} else {
							for (const param of apiCall.pathParams) {
								const value = document.getElementById('path-param-' + param).value;
								path = path.replace('{' + param + '}', value);
							}
							const paramsKeys = Object.keys(apiCall.params);
							if (paramsKeys.length > 0) {
								path += "?";
								paramsKeys.forEach((key, i) => {
									const value = document.getElementById('query-param-' + key).value;
									path += key + "=" + value;
									if (i !== paramsKeys.length - 1) {
										path += "&";
									}
								});
							}
						}
						localStorage.setItem("headers", JSON.stringify(headers));
						try {
							const res = await fetch(path, {
								method: apiCall.method,
								headers,
								...bodyParams
							});
							const json = await res.json();
							window.codeRight.setValue(JSON.stringify(json, 2, 2));
						} catch (err) {
							window.codeRight.setValue(JSON.stringify({
								error: err.message
							}, 2, 2));
						}
						run.innerHTML = "RUN";
					}
				</script>
			</body>
		</html>
	`).Props(
		"routes", apiRoutes,
		"apiData", template.HTML(string(apiData)),
	).Render()
}
