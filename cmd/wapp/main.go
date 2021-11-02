package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gobuffalo/velvet"
	"golang.org/x/mod/modfile"
)

var pathParamsRegex = regexp.MustCompile(`{(.*?)}`)

type Route struct {
	Method string
	Path   string
	Pkg    string
	Params []string
}

type ApiCall struct {
	Method string
	Name   string
	Params string
	Path   string
}

func getMethod(src string) string {
	if strings.HasSuffix(src, "get.go") {
		return "GET"
	} else if strings.HasSuffix(src, "post.go") {
		return "POST"
	} else if strings.HasSuffix(src, "put.go") {
		return "PUT"
	} else if strings.HasSuffix(src, "patch.go") {
		return "PATCH"
	} else if strings.HasSuffix(src, "delete.go") {
		return "DELETE"
	} else if strings.HasSuffix(src, "head.go") {
		return "HEAD"
	} else if strings.HasSuffix(src, "options.go") {
		return "OPTIONS"
	} else if strings.HasSuffix(src, "connect.go") {
		return "CONNECT"
	} else if strings.HasSuffix(src, "trace.go") {
		return "TRACE"
	} else {
		panic(fmt.Sprintf("Uknown route found %s", src))
	}
}

func getRoute(method, src string) string {
	return strings.ReplaceAll(src, "/"+strings.ToLower(method)+".go", "")
}

func getPackage(src string) string {
	return src
}

func rewritePath(route string) string {
	muxRoute := bytes.NewBuffer(nil)
	foundStart := false
	for _, v := range route {
		if string(v) == "_" && !foundStart {
			foundStart = true
			muxRoute.WriteString("{")
		} else if string(v) == "_" && foundStart {
			foundStart = false
			muxRoute.WriteString("}")
		} else {
			muxRoute.WriteString(string(v))
		}
	}
	return muxRoute.String()
}

func rewritePkg(pkg string) string {
	arr := strings.Split(pkg, "/")
	lastItem := arr[len(arr)-1]
	if strings.Contains(lastItem, "_") {
		return arr[len(arr)-2] + lastItem
	}
	return lastItem
}

func getApiFunc(method, route string, params []string) ApiCall {
	muxRoute := bytes.NewBuffer(nil)
	foundStart := false
	funcName := strings.ToLower(method)
	parts := strings.Split(route, "/")
	for _, p := range parts {
		if p != "api" {
			funcName += strings.Title(strings.Replace(p, "_", "", 2))
		}
	}
	for _, v := range route {
		if string(v) == "_" && !foundStart {
			foundStart = true
			muxRoute.WriteString("${")
		} else if string(v) == "_" && foundStart {
			foundStart = false
			muxRoute.WriteString("}")
		} else {
			muxRoute.WriteString(string(v))
		}
	}
	paramsStrings := ""
	if len(params) > 0 {
		paramsStrings += strings.Join(params, ", ") + ", params"
	}
	return ApiCall{
		Method: method,
		Name:   funcName,
		Params: paramsStrings,
		Path:   muxRoute.String(),
	}
}

// "io/ioutil"
// func migrate() {
// 	db := context.InitDB()
// 	ctx := c.Background()
// 	tx, err := context.BeginTransaction(db, ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	files, err := ioutil.ReadDir("./migrations")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, f := range files {
// 		data, err := ioutil.ReadFile("./migrations/" + f.Name())
// 		if err != nil {
// 			panic(err)
// 		}
// 		tx.MustExec(string(data))
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// "github.com/bxcodec/faker/v3"
// func seed() {
// 	db := context.InitDB()
// 	ctx := c.Background()
// 	tx, err := context.BeginTransaction(db, ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	reqContext := context.ReqContext{
// 		Tx:     tx,
// 		UserID: "123",
// 	}
// 	for i := 0; i < 20; i++ {
// 		ti := todos.TodoInput{}
// 		err := faker.FakeData(&ti)
// 		if err != nil {
// 			panic(err)
// 		}
// 		_, _, err = todos.POST(reqContext, ti)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	data, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("go.mod file not found %s", err.Error())
	}
	modTree, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatalf("could not parse go.mod %s", err.Error())
	}
	moduleName := modTree.Module.Mod.Path
	routes := []*Route{}
	apiCalls := []ApiCall{}
	allPkgs := map[string]string{}
	err = filepath.Walk("pages",
		func(filesrc string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				route := strings.ReplaceAll(filesrc, "pages", "")
				method := getMethod(route)
				path := getRoute(method, route)
				pkg := getPackage(path)
				allPkgs[pkg] = ""
				if path == "" { // for index page
					path = "/"
					pkg = "pages"
				}
				routePath := rewritePath(path)
				params := pathParamsRegex.FindAllString(routePath, -1)
				routes = append(routes, &Route{
					Method: method,
					Path:   routePath,
					Pkg:    rewritePkg(pkg),
					Params: params,
				})
				if strings.Contains(path, "/api/") {
					apiCalls = append(apiCalls, getApiFunc(method, path, params))
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range routes {
		fmt.Printf("%-6s %s\n", r.Method, r.Path)
	}
	ctx := velvet.NewContext()
	ctx.Set("moduleName", moduleName)
	ctx.Set("allPkgs", allPkgs)
	ctx.Set("routes", routes)
	ctx.Set("apiCalls", apiCalls)
	ctx.Set("tick", "`")
	s, err := velvet.Render(`// Code generated by wapp. DO NOT EDIT.
package main

import (
	"embed"
	"net/http"
	"os"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/mux"
	"github.com/pyros2097/wapp"
	"github.com/rs/zerolog/log"

	"{{ moduleName }}/context"
	{{#each allPkgs }}"{{ moduleName }}/pages{{ @key }}"
	{{/each}}
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	isLambda := os.Getenv("_LAMBDA_SERVER_PORT") != ""
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.PathPrefix("/assets/").Handler(http.FileServer(http.FS(assetsFS)))
	{{#each routes as |route| }}handle(r, "{{ route.Method }}", "{{ route.Path }}", {{ route.Pkg }}.{{ route.Method }})
	{{/each}}
	if !isLambda {
		println("http server listening on http://localhost:3000")
		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:3000",
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}
		log.Fatal().Stack().Err(srv.ListenAndServe()).Msg("failed to listen")
	} else {
		log.Print("running in lambda mode")
		log.Fatal().Stack().Err(gateway.ListenAndServe(":3000", r)).Msg("failed to listen")
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	wapp.LogReq(404, r)
}

func handle(router *mux.Router, method, route string, h interface{}) {
	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		var status int
		defer func() {
			wapp.LogReq(status, r)
		}()
		ctx, err := context.WithContext(r.Context())
		if err != nil {
			wapp.RespondError(w, 500, err)
			return
		}
		status, err = wapp.PerformRequest(route, h, ctx, w, r)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
		}
	}).Methods(method)
}
`, ctx)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("main.go", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
	js, err := velvet.Render(`// Code generated by wapp. DO NOT EDIT.
import queryString from 'query-string';
import config from './config';

const apiCall = async (method, route, params) => {
  const qs = method === 'GET' && params ? '?' + queryString.stringify(params) : '';
  const body = method !== 'GET' ? JSON.stringify(params) : null;
  const endpoint = await config.getApiEndpoint();
  const token = await config.getAuthToken();
  const res = await fetch({{tick}}${endpoint}/api/${route}${qs}{{tick}}, {
    method,
    headers: {
      Authorization: token,
    },
    body,
  });
  const json = await res.json();
  if (res.ok) {
    return json;
  } else {
    throw new Error(json.error);
  }
}

export default {
  {{#each apiCalls as |api| }}{{api.Name}}: ({{api.Params}}) => apiCall('{{api.Method}}', {{tick}}{{api.Path}}{{tick}}, params),
  {{/each}}
}
`, ctx)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("api.js", []byte(js), 0644)
	if err != nil {
		panic(err)
	}
}
