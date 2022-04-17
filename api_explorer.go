package gromer

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

	"github.com/carlmjohnson/versioninfo"
	"github.com/gorilla/mux"
	"github.com/pyros2097/gromer/assets"
	. "github.com/pyros2097/gromer/handlebars"
)

func ApiExplorerRoute(router *mux.Router, path string) {
	router.Path(path).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cmcss, _ := assets.FS.ReadFile("css/codemirror@5.63.1.css")
		stylescss, _ := assets.FS.ReadFile("css/styles.css")
		cmjs, _ := assets.FS.ReadFile("js/codemirror@5.63.1.min.js")
		cmjsjs, _ := assets.FS.ReadFile("js/codemirror-javascript@5.63.1.js")
		apiRoutes := []RouteDefinition{}
		for _, v := range RouteDefs {
			if strings.Contains(v.Path, "/api/") {
				apiRoutes = append(apiRoutes, v)
			}
		}
		apiData, err := json.Marshal(apiRoutes)
		if err != nil {
			RespondError(w, 500, err)
		}
		status, err := Html(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
				<meta http-equiv="encoding" content="utf-8">
				<title> API Explorer </title>
				<meta name="description" content="API Explorer">
				<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover">
				<style>
					{{ css }}
				</style>
				<script>
					{{ js }}
				</script>
			</head>
			<body>
				<div class="flex flex-col">
					<div class="flex w-full p-2 bg-gray-50 border-b border-gray-200 items-center justify-start">
						<div class="flex mr-4 text-gray-700 text-2xl font-bold"> API Explorer</div>
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
						<div class="flex flex-1 justify-end">
							(commit {{ commit }})
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
			"commit", versioninfo.Revision[0:7],
			"routes", apiRoutes,
			"apiData", template.HTML(string(apiData)),
			"css", template.HTML(string(cmcss)+"\n\n"+string(stylescss)),
			"js", template.HTML(string(cmjs)+"\n\n"+string(cmjsjs)),
		).RenderWriter(w)
		if err != nil {
			RespondError(w, status, err)
		}
	})
}
