package routes

import (
	"fmt"
	"net/http"
	"urlshortener/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Setup(handler *handlers.URLHandler) http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRD-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	mux.Use(middleware.Heartbeat(fmt.Sprintf("%s/ping", handler.Config.PathPrefix)))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.DefaultLogger)

	mux.Route(handler.Config.PathPrefix, func(mux chi.Router) {
		mux.Get("/", handler.HandleMainPage)
		mux.Get("/{key}", handler.HandleRedirectWithKey)
		mux.Post("/set", handler.HandleSetShortKey)
	})

	return mux
}
