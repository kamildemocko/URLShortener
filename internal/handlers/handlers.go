package handlers

import "urlshortener/internal/config"

type SetShortKeyRequestBody struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

type jsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Data    any    `json:"data,omitempty"`
}

type PageData struct {
	Protocol   string
	Domain     string
	SavedCount int
}

type URLHandler struct {
	Config *config.Config
}

func NewURLHandler(config *config.Config) *URLHandler {
	return &URLHandler{
		Config: config,
	}
}
