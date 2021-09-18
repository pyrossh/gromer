package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/mux"
	. "github.com/pyros2097/wapp"

	"github.com/pyros2097/wapp/example/context"
	"github.com/pyros2097/wapp/example/pages"
)

//go:embed assets/*
var assetsFS embed.FS

type Handler func(c *context.ReqContext) (interface{}, int, error)

func wrap(h Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewReqContext(w, r)
		value, status, err := h(ctx)
		w.WriteHeader(status)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			data, _ := json.Marshal(M{
				"error": err.Error(),
			})
			w.Write(data)
			return
		}
		if v, ok := value.(*Element); ok {
			w.Header().Set("Content-Type", "text/html")
			v.WriteHtml(w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(value)
		w.Write(data)
	}
}

func main() {
	isLambda := os.Getenv("_LAMBDA_SERVER_PORT") != ""
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.FileServer(http.FS(assetsFS)))
	r.HandleFunc("/", wrap(pages.Index))
	// r.HandleFunc("/about", wrap(pages.About))
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
