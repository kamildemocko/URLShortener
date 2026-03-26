package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"urlshortener/internal/config"
	"urlshortener/internal/handlers"
	"urlshortener/internal/repo/postgres"
	"urlshortener/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	config := config.Config{
		WebPort:    os.Getenv("PORT"),
		PathPrefix: "/",
		Dsn:        os.Getenv("DSN"),
	}

	repo, err := postgres.InitPostgresDB(config.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = repo.CreateTable()
	if err != nil {
		log.Fatal(err)
	}

	config.Repository = repo

	handlers := handlers.NewURLHandler(&config)
	routes := routes.Setup(handlers)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", config.WebPort),
		Handler:      routes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Starting server on port %s\n", config.WebPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
