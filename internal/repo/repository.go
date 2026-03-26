package repo

import (
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repository interface {
	CreateTable() error
	GetUrlByKey(key string) (string, error)
	SetKey(date time.Time, ip string, url string, key string) error
	KeyExists(key string) (bool, error)
	GetSavedCount() (int, error)
}

type QueryKeyEntry struct {
	Id        int
	Timestamp time.Time
	Ip        string
	Url       string
	Key       string
}

type QueryTableExists struct {
	Exists bool
}
