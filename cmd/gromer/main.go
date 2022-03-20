package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/gobuffalo/velvet"
	"github.com/pyros2097/gromer"
	"golang.org/x/mod/modfile"
)

type Route struct {
	Method string
	Path   string
	Pkg    string
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

func getApiFunc(method, route string, pathParams []string, params map[string]interface{}) gromer.ApiDefinition {
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
	return gromer.ApiDefinition{
		Method:     method,
		Path:       muxRoute.String(),
		PathParams: pathParams,
		Params:     params,
	}
}

func lowerFirst(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}
	return ""
}

func main() {
	moduleName := ""
	pkgFlag := flag.String("pkg", "", "specify a package name")
	flag.Parse()
	if pkgFlag == nil || *pkgFlag == "" {
		data, err := ioutil.ReadFile("go.mod")
		if err != nil {
			log.Fatalf("go.mod file not found %s", err.Error())
		}
		modTree, err := modfile.Parse("go.mod", data, nil)
		if err != nil {
			log.Fatalf("could not parse go.mod %s", err.Error())
		}
		moduleName = modTree.Module.Mod.Path
	} else {
		moduleName = *pkgFlag
	}
	routes := []*Route{}
	apiDefs := []gromer.ApiDefinition{}
	allPkgs := map[string]string{}
	err := filepath.Walk("pages",
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
				pathParams := gromer.GetRouteParams(routePath)
				routes = append(routes, &Route{
					Method: method,
					Path:   routePath,
					Pkg:    rewritePkg(pkg),
				})
				if strings.Contains(path, "/api/") {
					data, err := ioutil.ReadFile(filesrc)
					if err != nil {
						panic(err)
					}
					fset := token.NewFileSet()
					f, err := parser.ParseFile(fset, "", string(data), parser.AllErrors)
					if err != nil {
						panic(err)
					}
					var params map[string]interface{}
					mapsOfInputParams := map[string]map[string]interface{}{}
					for _, d := range f.Decls {
						if decl, ok := d.(*ast.GenDecl); ok {
							for _, spec := range decl.Specs {
								if typeSpec, ok := spec.(*ast.TypeSpec); ok {
									if st, ok := typeSpec.Type.(*ast.StructType); ok {
										mapsOfInputParams[typeSpec.Name.Name] = map[string]interface{}{}
										for _, f := range st.Fields.List {
											if len(f.Names) > 0 {
												fieldName := lowerFirst(f.Names[0].Name)
												mapsOfInputParams[typeSpec.Name.Name][fieldName] = fmt.Sprintf("%+s", f.Type)
											}
										}
									}
								}
							}
						}
						if decl, ok := d.(*ast.FuncDecl); ok {
							if decl.Name.Name == method {
								list := decl.Type.Params.List
								lastParam := list[len(list)-1]
								if spec, ok := lastParam.Type.(*ast.Ident); ok {
									if spec.IsExported() {
										// TODO: need to read from imported files and pick up the struct and
										// convert it to json
									} else if v, ok := mapsOfInputParams[spec.Name]; ok {
										params = v
									}
								}
							}
						}
					}
					apiDefs = append(apiDefs, getApiFunc(method, path, pathParams, params))
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
	ctx.Set("apiDefs", apiDefs)
	ctx.Set("tick", "`")
	s, err := velvet.Render(`// Code generated by gromer. DO NOT EDIT.
package main

import (
	c "context"
	"embed"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pyros2097/gromer"
	"github.com/rs/zerolog/log"
	"gocloud.dev/server"

	"{{ moduleName }}/context"
	{{#each allPkgs }}"{{ moduleName }}/pages{{ @key }}"
	{{/each}}
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.PathPrefix("/assets/").Handler(wrapCache(http.FileServer(http.FS(assetsFS))))
	handle(r, "GET", "/api", gromer.ApiExplorer(apiDefinitions()))
	{{#each routes as |route| }}handle(r, "{{ route.Method }}", "{{ route.Path }}", {{ route.Pkg }}.{{ route.Method }})
	{{/each}}
	println("http server listening on http://localhost:"+port)
	srv := server.New(r, nil)
	if err := srv.ListenAndServe(":"+port); err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to listen")
	}
}

func wrapCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=2592000")
		h.ServeHTTP(w, r)
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	gromer.LogReq(404, r)
}

func handle(router *mux.Router, method, route string, h interface{}) {
	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		var status int
		defer func() {
			gromer.LogReq(status, r)
		}()
		ctx := c.WithValue(
			c.WithValue(
				c.WithValue(r.Context(), "assetsFS", assetsFS),
					"url", r.URL),
			"header", r.Header)
		status, err = gromer.PerformRequest(route, h, ctx, w, r)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
		}
	}).Methods(method)
}

func apiDefinitions() []gromer.ApiDefinition {
	return []gromer.ApiDefinition{
		{{#each apiDefs as |api| }}
		{
			Method: "{{api.Method}}",
			Path: "{{api.Path}}",
			PathParams: []string{ {{#each api.PathParams as |param| }}"{{param}}", {{/each}} },
			Params: map[string]interface{}{
				{{#each api.Params }}"{{ @key }}": "{{ @value }}", {{/each}}
			},
		},{{/each}}
	}
}
`, ctx)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("main.go", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
