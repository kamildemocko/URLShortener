package main

import (
	"fmt"
	"net/http"
	"strings"
)

const allowedUrlCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~:/?#[]@!$&'()*+,;="

func (app *Config) ValidateUrl(url string) error {
	if !strings.HasPrefix(url, "http") {
		return fmt.Errorf("url has to start with http or https at the beginning")
	}

	if len(url) > 2048 {
		return fmt.Errorf("url is too long")
	}

	if !strings.ContainsAny(url, (allowedUrlCharacters)) {
		return fmt.Errorf("key contains unsupported character(s)")
	}

	// test url if it works -- for redirect to work it has to start with http*
	_, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("input URL doesn't seem to work")
	}

	return nil
}

func (app *Config) ValidateKey(key string) error {
	if len(key) < 2 || len(key) > 32 {
		return fmt.Errorf("key has to be between 2 - 32 characters")
	}

	if !strings.ContainsAny(key, (allowedUrlCharacters)) {
		return fmt.Errorf("key contains unsupported character(s)")
	}

	return nil
}
