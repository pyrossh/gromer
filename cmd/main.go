package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/velvet"
	"golang.org/x/mod/modfile"
)

type Route struct {
	Method string
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
	muxRoute := bytes.NewBuffer(nil)
	baseRoute := strings.Replace(src, "/"+strings.ToLower(method)+".go", "", 1)
	foundStart := false
	for _, v := range baseRoute {
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

func main() {
	data, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("go.mod file not found %w", err)
	}
	modTree, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatalf("could not parse go.mod %w", err)
	}
	moduleName := modTree.Module.Mod.Path
	routes := []*Route{}
	err = filepath.Walk("pages",
		func(filesrc string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				route := strings.Replace(filesrc, "pages", "", 1)
				method := getMethod(route)
				path := getRoute(method, route)
				if path == "" { // for index page
					path = "/"
				}
				routes = append(routes, &Route{Method: method, Path: path})
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range routes {
		println(r.Method, r.Path)
	}
	ctx := velvet.NewContext()
	ctx.Set("moduleName", moduleName)
	ctx.Set("routes", routes)
	s, err := velvet.Render(`// GENERATED BY WAPP DO NOT EDIT THIS FILE
package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/mux"
	. "{{ moduleName }}/context"
	"{{ moduleName }}/pages"
	"{{ moduleName }}/pages/about"
	"{{ moduleName }}/pages/api/todos/_id_"
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	isLambda := os.Getenv("_LAMBDA_SERVER_PORT") != ""
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.FileServer(http.FS(assetsFS)))

	{{#each routes as |route| }}r.HandleFunc("{{ route.Path }}", Wrap(pages.GET)).Methods("{{ route.Method }}")
	{{/each}}
	if !isLambda {
		println("http server listening on http://localhost:3000")
		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:3000",
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
	} else {
		log.Print("running in lambda mode")
		log.Fatal(gateway.ListenAndServe(":3000", r))
	}
}
	
	`, ctx)
	if err != nil {
		panic(err)
	}
	println(s)
}
