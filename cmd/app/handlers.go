package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type PageData struct {
	Protocol string
	Domain   string
}

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
	if len(ip) > 32 {
		ip = ip[:32]
	}

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

	newUrl := fmt.Sprintf("%s://%s/short/%s", os.Getenv("PROTOCOL"), os.Getenv("DOMAIN"), inputRequest.Key)
	app.WriteJSON(w, http.StatusOK, "success", newUrl, nil)
}

func (app *Config) handleMainPage(w http.ResponseWriter, r *http.Request) {
	render(w, "main.page.gohtml", PageData{Protocol: os.Getenv("PROTOCOL"), Domain: os.Getenv("DOMAIN")})
}

func render(w http.ResponseWriter, t string, data PageData) {
	partials := []string{
		"./templates/base.layout.gohtml",
		"./templates/header.partial.gohtml",
		"./templates/footer.partial.gohtml",
		"./templates/modalerror.partial.gohtml",
		"./templates/modalsuccess.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
