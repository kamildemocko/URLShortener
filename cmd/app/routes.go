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
		AllowedMethods: []string{"GET", "PUT"},
		AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRD-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	mux.Use(middleware.Heartbeat(fmt.Sprintf("%s/ping", app.pathPrefix)))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.DefaultLogger)

	mux.Handle("/favicon.ico", http.FileServer(http.Dir("./")))

	mux.Route(app.pathPrefix, func(mux chi.Router) {
		mux.HandleFunc("GET /go/{key}", app.handleRedirectWithKey)
		mux.HandleFunc("PUT /set", app.handleSetShortKey)
		mux.HandleFunc("GET /", app.handleMainPage)
	})

	return mux
}
