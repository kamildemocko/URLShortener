package main

import (
	"net/http"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", getUrlByShortName)

	wrappedMux := wrapMiddlewares(
		mux,
		recoverer,
		ping,
	)

	return wrappedMux
}
