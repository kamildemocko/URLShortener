package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func wrapMiddlewares(handler http.Handler, mdws ...func(http.Handler) http.Handler) http.Handler {
	for _, mdw := range mdws {
		handler = mdw(handler)
	}

	return handler
}

func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n%s", err, debug.Stack())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func ping(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ping" {
			w.Write([]byte("pong"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
