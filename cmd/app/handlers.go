package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) handleRedirectWithKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	fmt.Println(key)

	// get url from db
	url := "https://google.com"

	// redirect to url
	http.Redirect(w, r, url, http.StatusFound)
}

func (app *Config) handleSetShortKey(w http.ResponseWriter, r *http.Request) {
	var inputRequest SetShortKeyRequestBody

	err := json.NewDecoder(r.Body).Decode(&inputRequest)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("cannot parse request body json"))
	}
}
