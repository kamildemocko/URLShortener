package main

import "net/http"

func handleShortKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi!"))
}

func handleSetShortKey(w http.ResponseWriter, r *http.Request) {
}
