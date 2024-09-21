package main

import (
	"net/http"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/go/{shortKey}", handleShortKey)
	mux.HandleFunc("/set/{shortKey}", handleSetShortKey)

	wrappedMux := wrapMiddlewares(
		mux,
		recoverer,
		ping,
	)

	return wrappedMux
}
