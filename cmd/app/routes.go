package main

import (
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

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.DefaultLogger)

	mux.Route("/short", func(mux chi.Router) {
		mux.HandleFunc("/go/{shortKey}", app.handleShortKey)
		mux.HandleFunc("/set", app.handleSetShortKey)
	})

	return mux
}
