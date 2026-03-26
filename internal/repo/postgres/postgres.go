package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"urlshortener/internal/repo"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type postgresRepository struct {
	DB *sql.DB
}

func InitPostgresDB(dsn string) (repo.Repository, error) {
	log.Println("connecting to DB")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(5 * time.Minute)

	repo := &postgresRepository{
		DB: db,
	}

	return repo, nil
}

func (pr *postgresRepository) CreateTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queryTableExists := `
		SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'urlshortener' 
		AND table_name = 'keys'
		);`

	queryTable := `
		CREATE TABLE IF NOT EXISTS urlshortener.keys (
		id SERIAL PRIMARY KEY,
		timestamp TIMESTAMP WITH TIME ZONE,
		ip VARCHAR(32),
		url VARCHAR(2048),
		key VARCHAR(32),
		CONSTRAINT unique_key UNIQUE (key)
		);`

	queryIndex := `
		CREATE INDEX idx_keys_key ON urlshortener.keys(key)`

	var tableExists repo.QueryTableExists
	pr.DB.QueryRowContext(ctx, queryTableExists).Scan(&tableExists.Exists)
	if tableExists.Exists {
		return nil
	}

	tx, err := pr.DB.Begin()
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, queryTable); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.ExecContext(ctx, queryIndex); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (pr *postgresRepository) GetUrlByKey(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT * FROM urlshortener.keys
		WHERE key = $1;`

	row := pr.DB.QueryRowContext(ctx, query, key)

	var data repo.QueryKeyEntry
	if err := row.Scan(
		&data.Id,
		&data.Timestamp,
		&data.Ip,
		&data.Url,
		&data.Key,
	); err != nil {
		return "", err
	}

	return data.Url, nil
}

func (pr *postgresRepository) SetKey(date time.Time, ip string, url string, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO urlshortener.keys (timestamp, ip, url, key) 
		VALUES ($1, $2, $3, $4);`

	_, err := pr.DB.ExecContext(ctx, query, date, ip, url, key)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pr *postgresRepository) GetSavedCount() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT COUNT(*) FROM urlshortener.keys;`

	row := pr.DB.QueryRowContext(ctx, query)

	var count int
	if err := row.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}
