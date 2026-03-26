package handlers

import (
	"net/http"
	"urlshortener/internal/components"
)

func (u *URLHandler) HandleMainPage(w http.ResponseWriter, r *http.Request) {
	savedCount, err := u.Config.Repository.GetSavedCount()
	if err != nil {
		u.HandleInternalErrorPage(w, r)
		return
	}

	component := components.AddFrame("Make Short", savedCount)
	_ = component.Render(r.Context(), w)
}

func (u *URLHandler) HandleSuccessPage(w http.ResponseWriter, r *http.Request, url string) {
	w.WriteHeader(http.StatusOK)
	component := components.ShowFrame(url)
	_ = component.Render(r.Context(), w)
}

func (u *URLHandler) HandleErrorPage(w http.ResponseWriter, r *http.Request, err string) {
	w.WriteHeader(http.StatusNotFound)
	component := components.ErrorFrame(err)
	_ = component.Render(r.Context(), w)
}

func (u *URLHandler) HandleNotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	component := components.NotExistsFrame()
	_ = component.Render(r.Context(), w)
}

func (u *URLHandler) HandleInternalErrorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	component := components.InternalErrorFrame()
	_ = component.Render(r.Context(), w)
}
