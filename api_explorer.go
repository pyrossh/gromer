package gromer

import (
	"context"
	"encoding/json"
	"fmt"
)

func Select(uis ...interface{}) *Element {
	return NewElement("select", false, uis...)
}

func Option(uis ...interface{}) *Element {
	return NewElement("option", false, uis...)
}

func Table(uis ...interface{}) *Element {
	return NewElement("table", false, uis...)
}

func TR(uis ...interface{}) *Element {
	return NewElement("tr", false, uis...)
}

func TD(uis ...interface{}) *Element {
	return NewElement("td", false, uis...)
}

func Section(title string) *Element {
	return Div(Css("text-gray-700 text-sm font-bold uppercase pl-2 pt-2 pb-2 bg-gray-50 border-b border-gray-200"), Text(title))
}

type ApiDefinition struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	PathParams  []string          `json:"pathParams"`
	QueryParams map[string]string `json:"queryParams"`
}

func ApiExplorer(apiDefs []ApiDefinition) func(c context.Context) (HtmlPage, int, error) {
	return func(c context.Context) (HtmlPage, int, error) {
		data, err := json.Marshal(apiDefs)
		if err != nil {
			return Html(nil, nil), 500, err
		}
		options := []interface{}{ID("api-select"), Css("form-select block")}
		for i, c := range apiDefs {
			options = append(options, Option(Attr("value", fmt.Sprintf("%d", i)), Div(Text(fmt.Sprintf("%0.6s %s", c.Method, c.Path)))))
		}
		return Html(
			Head(
				Title("Example"),
				Meta("description", "Example"),
				Meta("author", "pyros.sh"),
				Meta("keywords", "pyros.sh, gromer"),
				Meta("viewport", "width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover"),
				Link("icon", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQBAMAAADt3eJSAAAAMFBMVEU0OkArMjhobHEoPUPFEBIuO0L+AAC2FBZ2JyuNICOfGx7xAwTjCAlCNTvVDA1aLzQ3COjMAAAAVUlEQVQI12NgwAaCDSA0888GCItjn0szWGBJTVoGSCjWs8TleQCQYV95evdxkFT8Kpe0PLDi5WfKd4LUsN5zS1sKFolt8bwAZrCaGqNYJAgFDEpQAAAzmxafI4vZWwAAAABJRU5ErkJggg=="),
				Link("stylesheet", "https://cdn.jsdelivr.net/npm/codemirror@5.63.1/lib/codemirror.css"),
				StyleTag(Text(`
					html, body {
						height: 100vh;
					} 
	
					#left .CodeMirror {
						height: 400px; 
					}

					#right .CodeMirror {
						height: calc(100vh - 60px); 
					}
	
					.form-select {
						background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23a0aec0'%3e%3cpath d='M15.3 9.3a1 1 0 0 1 1.4 1.4l-4 4a1 1 0 0 1-1.4 0l-4-4a1 1 0 0 1 1.4-1.4l3.3 3.29 3.3-3.3z'/%3e%3c/svg%3e");
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
	
					textarea:focus, input:focus{
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
				`)),
				Script(Src("https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.63.1/codemirror.min.js")),
				Script(Src("https://cdn.jsdelivr.net/npm/codemirror@5.63.1/mode/javascript/javascript.js")),
			),
			Body(
				Div(Css("flex flex-col"),
					Div(Css("flex w-full p-2 bg-gray-50 border-b border-gray-200 items-center justify-start"),
						Div(Css("flex mr-4 text-gray-700 text-2xl font-bold"), Text("API Explorer")),
						Div(Css("text-xl"),
							Select(options...),
						),
						Div(Css("flex ml-3 mr-3"), Button(ID("run"), Css("bg-gray-200 border border-gray-400 hover:bg-gray-200 focus:outline-none rounded-md text-gray-700 text-md font-bold pt-2 pb-2 pl-6 pr-6"),
							Text("RUN"),
						)),
					),
					Div(Css("flex"),
						Div(Css("flex flex-row"), Style("width: 50%;"),
							Div(Css("pr-8 border-r border-gray-300"), Style("background: #f7f7f7;")),
							Div(Css("w-full"),
								Section("Headers"),
								Table(ID("headersTable"),
									TR(
										TD(Input(Css("w-full p-1"), Attr("value", "Authorization"))),
										TD(Input(Css("w-full p-1"))),
									),
								),
								Section("Path Params"),
								Table(ID("pathParamsTable"),
									TR(
										TD(Css("text-gray-700"), Style("width: 50%;"), Div(Css("p-1"), Text("123"))),
										TD(Style("width: 50%;"), Input(Css("w-full p-1"))),
									),
								),
								Section("Query Params"),
								Table(ID("queryParamsTable"),
									TR(
										TD(Css("text-gray-700"), Style("width: 50%;"), Div(Css("p-1"), Text("123"))),
										TD(Style("width: 50%;"), Input(Css("w-full p-1"))),
									),
								),
								Section("Body"),
								Div(ID("left"), Css("border-b border-gray-200 text-md")),
							),
						),
						Div(Css("flex flex-row"), Style("width: 50%;"),
							Div(ID("right"), Css("w-full border-l border-l-gray-200 text-md")),
							Div(Css("pr-8 border-l border-gray-300"), Style("background: #f7f7f7;")),
						),
					),
				),
				Script(Text("window.apiDefs = "+string(data))),
				Script(Text("window.codeLeft = CodeMirror(document.getElementById('left'), {value: '{}',mode:  'javascript' })")),
				Script(Text("window.codeRight = CodeMirror(document.getElementById('right'), {value: '',mode:  'javascript', lineNumbers: true, readOnly: true, lineWrapping: true })")),
				Script(Text(`
					const getCurrentApiCall = () => {
						const index = document.getElementById("api-select").value;
						return window.apiDefs[index];
					}
	
					const updatePathParams = (apiCall) => {
						const table = document.getElementById("pathParamsTable");
						if (apiCall.pathParams.length === 0) {
							table.innerHTML = "<div style='background-color: rgb(245, 245, 245); padding: 0.25rem; text-align: center; color: gray;'>NONE</div>";
						} else {
							table.innerHTML = "";
						}
						for(const param of apiCall.pathParams.reverse()) {
							const row = table.insertRow(0);
							const cell1 = row.insertCell(0);
							const cell2 = row.insertCell(1);
							cell1.style = "width: 30%; border-left: 0px;";
							cell1.class = "text-gray-700";
							cell2.style = "width: 70%;";
							cell1.innerHTML = "<div class='p-1'>" + param + "</div>";
							cell2.innerHTML = "<input id='path-param-" + param + "' class='w-full p-1'>";
						}
					}
	
					const updateQueryParams = (apiCall) => {
						const table = document.getElementById("queryParamsTable");
						if (!apiCall.queryParams) {
							table.innerHTML = "<div style='background-color: rgb(245, 245, 245); padding: 0.25rem; text-align: center; color: gray;'>NONE</div>";
						} else {
							table.innerHTML = "";
						}
					}
	
					const updateBody = (apiCall) => {
						const editor = document.getElementById("left");
					}
					
					const init = () => {
						updatePathParams(window.apiDefs[0]);
						updateQueryParams(window.apiDefs[0]);
						const headersJson = localStorage.getItem("headers");
						if (headersJson) {
							const table = document.getElementById("headersTable");
							const headers = JSON.parse(headersJson);
							table.innerHTML = "";
							for(const key of Object.keys(headers)) {
								const value = headers[key];
								const row = table.insertRow(0);
								const cell1 = row.insertCell(0);
								const cell2 = row.insertCell(1);
								cell1.style = "width: 30%; border-left: 0px;";
								cell2.style = "width: 70%;";
								cell1.innerHTML = "<input value='" + key + "' class='w-full p-1'>";
								cell2.innerHTML = "<input value='" + value + "' class='w-full p-1'>";
							}
						}
					}
	
					window.onload = () => {
						init();
					}
	
					document.getElementById("api-select").onchange = () => {
						const apiCall = getCurrentApiCall();
						updatePathParams(apiCall);
						updateQueryParams(apiCall);
						updateBody(apiCall);
					}
					
					const run = document.getElementById("run");
					run.onclick = async () => {
						run.innerHTML = "<svg class='spinner' viewBox='0 0 50 50'><circle class='path' cx='25' cy='25' r='20' fill='none' stroke-width='5'></circle></svg>";
						const table = document.getElementById("headersTable");
						const headers = {};
						for(const row of table.rows) {
							const key = row.cells[0].children[0].value;
							const value = row.cells[1].children[0].value;
							headers[key] = value;
						}
						const apiCall = getCurrentApiCall();
						let path = apiCall.path;
						const bodyParams = {};
						if (apiCall.method !== "GET" && apiCall.method != "DELETE") {
							bodyParams["body"] = window.codeLeft.getValue();
						}
						for(const param of apiCall.pathParams) {
							const value = document.getElementById('path-param-' + param).value;
							path = path.replace('{' + param + '}', value);
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
							window.codeRight.setValue(JSON.stringify({ error: err.message }, 2, 2));
						}
						run.innerHTML = "RUN";
					}
				`)),
			),
		), 200, nil
	}
}
