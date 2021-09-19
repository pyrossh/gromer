package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/apex/gateway/v2"
	"github.com/gorilla/mux"
	. "github.com/pyros2097/wapp/example/context"
	"github.com/pyros2097/wapp/example/pages"
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	isLambda := os.Getenv("_LAMBDA_SERVER_PORT") != ""
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.FileServer(http.FS(assetsFS)))
	r.HandleFunc("/", Wrap(pages.Index)).Methods("GET")
	r.HandleFunc("/about", Wrap(pages.About)).Methods("GET")
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
