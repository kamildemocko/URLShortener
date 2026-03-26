package config

import "urlshortener/internal/repo"

type Config struct {
	WebPort    string
	PathPrefix string
	Dsn        string
	Repository repo.Repository
}
