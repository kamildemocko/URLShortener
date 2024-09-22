package data

import (
	"time"
)

type modKeysKey struct {
	id        int
	timestamp time.Time
	ip        string
	url       string
	key       string
}

type modTableKeysExists struct {
	exists bool
}
