package main

import (
	"URLShortener/data"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func initPostgresDB(dsn string) (data.Repository, error) {
	log.Println("connecting to DB")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(5 * time.Minute)

	repo := data.NewPostgresDB(db)
	if err = repo.CreateTable(); err != nil {
		return nil, err
	}

	return repo, nil
}
