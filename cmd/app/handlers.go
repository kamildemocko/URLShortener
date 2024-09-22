package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) handleRedirectWithKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	url, err := app.repository.GetUrlByKey(key)
	if err != nil {
		_ = app.NotFoundJSON(w, fmt.Errorf("cannot find url by key: %s", key))
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func (app *Config) handleSetShortKey(w http.ResponseWriter, r *http.Request) {
	var inputRequest SetShortKeyRequestBody

	err := json.NewDecoder(r.Body).Decode(&inputRequest)
	if err != nil {
		_ = app.ErrorJSON(w, fmt.Errorf("cannot parse request body json"))
		return
	}
}
