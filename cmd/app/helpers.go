package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

type jsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) WriteJSON(w http.ResponseWriter, status int, msg, detail string, data any) error {
	payloadData := jsonResponse{
		Code:    status,
		Message: msg,
		Detail:  detail,
		Data:    data,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(payload)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) ErrorJSON(w http.ResponseWriter, err error) error {
	statusCode := http.StatusBadRequest

	return app.WriteJSON(w, statusCode, "error", err.Error(), nil)
}

func (app *Config) NotFoundJSON(w http.ResponseWriter, err error) error {
	statusCode := http.StatusNotFound

	return app.WriteJSON(w, statusCode, "error", err.Error(), nil)
}

func (app *Config) AlreadyExistsJSON(w http.ResponseWriter, err error) error {
	statusCode := http.StatusConflict

	return app.WriteJSON(w, statusCode, "error", err.Error(), nil)
}

func (app *Config) GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	if ip == "" {
		ip = r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip, _, _ = net.SplitHostPort(ip)
		}
	}
	return ip
}
