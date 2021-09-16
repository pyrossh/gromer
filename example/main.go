package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/mux"
	. "github.com/pyros2097/wapp"
	"github.com/pyros2097/wapp/example/pages"
)

//go:embed assets/*
var assetsFS embed.FS

func wrap(f func(http.ResponseWriter, *http.Request) *Element) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		f(w, r).WriteHtml(w)
	}
}

func main() {
	isLambda := os.Getenv("_LAMBDA_SERVER_PORT") != ""
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.FileServer(http.FS(assetsFS)))
	r.HandleFunc("/", wrap(pages.Index))
	r.HandleFunc("/about", wrap(pages.About))
	if !isLambda {
		println("http listening on http://localhost:1234")
		srv := &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:1234",
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}
		log.Fatal(srv.ListenAndServe())
	} else {
		log.Print("running in lambda mode")
		log.Fatal(gateway.ListenAndServe(":3000", r))
	}
}
