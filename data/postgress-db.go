package data

import (
	"context"
	"time"
)

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
		ip VARCHAR(16),
		url VARCHAR(2048),
		key VARCHAR(32),
		CONSTRAINT unique_key UNIQUE (key)
		);`

	queryIndex := `
		CREATE INDEX idx_keys_key ON urlshortener.keys(key)`

	var tableExists modTableKeysExists
	pr.DB.QueryRowContext(ctx, queryTableExists).Scan(&tableExists.exists)
	if tableExists.exists {
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

	if err = tx.Commit(); err != nil {
		return nil
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

	var data modKeysKey
	if err := row.Scan(
		&data.id,
		&data.timestamp,
		&data.ip,
		&data.url,
		&data.key,
	); err != nil {
		return "", err
	}

	return data.url, nil
}

// func (pr *postgresRepository) SetKey(string, string, string) error {}
