package main

import "net/http"

func getUrlByShortName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi!"))
}

func setUrlShortName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi!"))
}
