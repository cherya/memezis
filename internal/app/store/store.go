package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Store struct {
	db      *sqlx.DB
	baseURL string
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

var ErrNotFound = errors.New("not found")
