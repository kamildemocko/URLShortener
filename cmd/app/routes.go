package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRD-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	mux.Use(middleware.Heartbeat(fmt.Sprintf("%s/ping", app.pathPrefix)))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.DefaultLogger)

	mux.Route(app.pathPrefix, func(mux chi.Router) {
		mux.HandleFunc("/go/{key}", app.handleRedirectWithKey)
		mux.HandleFunc("/set", app.handleSetShortKey)
	})

	return mux
}
