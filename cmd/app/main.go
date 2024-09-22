package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	webPort    int
	pathPrefix string
	dsn        string
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf(".env file cannot be loaded, detail: %s", err.Error()))
	}
}

func main() {
	app := Config{
		webPort:    80,
		pathPrefix: "/short",
		dsn:        os.Getenv("DSN"),
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.webPort),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Starting server on port %d\n", app.webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
