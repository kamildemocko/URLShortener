package main

import (
	"URLShortener/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	webPort    string
	pathPrefix string
	dsn        string
	repository data.Repository
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf(".env file cannot be loaded, detail: %s", err.Error()))
	}
}

func main() {
	app := Config{
		webPort:    os.Getenv("PORT"),
		pathPrefix: "/short",
		dsn:        os.Getenv("DSN"),
	}

	repo, err := initPostgresDB(app.dsn)
	if err != nil {
		panic(err)
	}
	app.repository = repo

	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", app.webPort),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Starting server on port %s\n", app.webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
