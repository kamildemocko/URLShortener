package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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

	if inputRequest.URL == "" || inputRequest.Key == "" {
		app.ErrorJSON(w, fmt.Errorf("missing 'url' or 'key' parameter"))
		return
	}

	if err = app.ValidateUrl(inputRequest.URL); err != nil {
		app.ErrorJSON(w, err)
		return
	}

	if err = app.ValidateKey(inputRequest.Key); err != nil {
		app.ErrorJSON(w, err)
		return
	}

	ip := app.GetIP(r)

	err = app.repository.SetKey(time.Now(), ip, inputRequest.URL, inputRequest.Key)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "unique_key"):
			app.AlreadyExistsJSON(w, fmt.Errorf("key already exists"))
		default:
			app.ErrorJSON(w, fmt.Errorf("cannot save key at this time"))
			log.Println(err.Error())
		}

		return
	}

	app.WriteJSON(w, http.StatusOK, "success", "data inserted", nil)
}
