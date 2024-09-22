package data

import "database/sql"

type Repository interface {
	CreateTable() error
	GetUrlByKey(string) (string, error)
	// SetKey(string, string, string) error
}

type postgresRepository struct {
	DB *sql.DB
}

func NewPostgresDB(db *sql.DB) Repository {
	return &postgresRepository{
		DB: db,
	}
}
