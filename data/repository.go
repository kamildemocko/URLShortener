package data

import (
	"database/sql"
	"time"
)

type Repository interface {
	CreateTable() error
	GetUrlByKey(key string) (string, error)
	SetKey(date time.Time, ip string, url string, key string) error
}

type postgresRepository struct {
	DB *sql.DB
}

func NewPostgresDB(db *sql.DB) Repository {
	return &postgresRepository{
		DB: db,
	}
}
