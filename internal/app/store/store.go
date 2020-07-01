package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type store struct {
	db      *sqlx.DB
	baseURL string
}

func NewStore(db *sqlx.DB) *store {
	return &store{
		db: db,
	}
}

var ErrNotFound = errors.New("not found")
