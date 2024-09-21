package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const webPort = 80

func main() {
	mux := newMux()

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", webPort),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Printf("Starting server on port %d\n", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

}
