package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (u *URLHandler) HandleRedirectWithKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	url, err := u.Config.Repository.GetUrlByKey(key)
	if err != nil {
		http.Redirect(w, r, "/notfound", http.StatusSeeOther)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}
