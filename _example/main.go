package main

import (
	"github.com/gorilla/mux"
	"github.com/pyros2097/gromer"
	gromer_assets "github.com/pyros2097/gromer/assets"
	"github.com/pyros2097/gromer/gsx"
	"github.com/rs/zerolog/log"
	"gocloud.dev/server"

	"github.com/pyros2097/gromer/_example/assets"
	"github.com/pyros2097/gromer/_example/components"
	"github.com/pyros2097/gromer/_example/containers"
	"github.com/pyros2097/gromer/_example/routes"
)

func init() {
	gsx.RegisterComponent(components.Todo, components.TodoStyles, "todo")
	gsx.RegisterComponent(components.Status, components.StatusStyles, "status", "error")
	gsx.RegisterComponent(containers.TodoCount, nil, "filter")
	gsx.RegisterComponent(containers.TodoList, nil, "page", "filter")
}

func main() {
	baseRouter := mux.NewRouter()
	baseRouter.Use(gromer.LogMiddleware)
	gromer.RegisterStatusHandler(baseRouter, components.Status)

	staticRouter := baseRouter.NewRoute().Subrouter()
	staticRouter.Use(gromer.CacheMiddleware)
	staticRouter.Use(gromer.CompressMiddleware)
	gromer.StaticRoute(staticRouter, "/gromer/", gromer_assets.FS)
	gromer.StaticRoute(staticRouter, "/assets/", assets.FS)
	gromer.IconsRoute(staticRouter, "/icons/", assets.FS)
	gromer.ComponentStylesRoute(staticRouter, "/components.css")

	pageRouter := baseRouter.NewRoute().Subrouter()
	gromer.PageRoute(pageRouter, "/", routes.TodosPage, routes.TodosAction)
	gromer.PageRoute(pageRouter, "/about", routes.AboutPage, nil)

	log.Info().Msg("http server listening on http://localhost:3000")
	srv := server.New(baseRouter, nil)
	if err := srv.ListenAndServe(":3000"); err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to listen")
	}
}
