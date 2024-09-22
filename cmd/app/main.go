package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const webPort = 80

type Config struct {
	Mux http.Handler
}

func main() {
	app := Config{}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", webPort),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Starting server on port %d\n", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
