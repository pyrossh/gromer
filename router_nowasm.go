// +build !wasm

package app

import (
	"bytes"
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pyros2097/wapp/errors"
)

// ServeHTTP makes the router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Handle errors
	defer func() {
		if rcv := recover(); rcv != nil {
			var err error
			switch x := rcv.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			writePage(r.Error(NewRenderContext(), err), w)
		}
	}()

	path := req.URL.Path
	println("route: " + req.URL.Path)

	if root := r.trees[req.Method]; root != nil {
		// TODO: use _ ps save it to context for useParam()
		if handle, _, tsr := root.getValue(path, nil); handle != nil {
			if render, ok := handle.(RenderFunc); ok {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(http.StatusOK)
				writePage(render(NewRenderContext()), w)
				return
			} else {
				handle.(func(w http.ResponseWriter, r *http.Request))(w, req)
				return
			}

		} else if req.Method != http.MethodConnect && path != "/" {
			// Moved Permanently, request with GET method
			code := http.StatusMovedPermanently
			if req.Method != http.MethodGet {
				// Permanent Redirect, request with same method
				code = http.StatusPermanentRedirect
			}

			if tsr && r.RedirectTrailingSlash {
				if len(path) > 1 && path[len(path)-1] == '/' {
					req.URL.Path = path[:len(path)-1]
				} else {
					req.URL.Path = path + "/"
				}
				http.Redirect(w, req, req.URL.String(), code)
				return
			}
		}
	}

	// Handle 404
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	writePage(r.NotFound(NewRenderContext()), w)
}

func (r *Router) getPage(ui UI) string {
	b := bytes.NewBuffer(nil)
	writePage(ui, b)
	return b.String()
}

func (r *Router) Lambda(ctx context.Context, e events.APIGatewayV2HTTPRequest) (res events.APIGatewayV2HTTPResponse, err error) {
	res.StatusCode = 200
	res.Headers = map[string]string{
		"Content-Type": "text/html",
	}
	// Handle errors
	defer func() {
		if rcv := recover(); rcv != nil {
			switch x := rcv.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
			res.Body = r.getPage(r.Error(NewRenderContext(), err))
		}
	}()

	println("route: " + e.RawPath)
	path := strings.Replace(e.RawPath, "/Prod/", "/", 1)
	if root := r.trees[e.RequestContext.HTTP.Method]; root != nil {
		if handle, _, _ := root.getValue(path, nil); handle != nil {
			res.Body = r.getPage(handle.(RenderFunc)(NewRenderContext()))
			return
		}
	}

	// Handle 404
	res.StatusCode = http.StatusNotFound
	res.Body = r.getPage(r.NotFound(NewRenderContext()))
	return
}
