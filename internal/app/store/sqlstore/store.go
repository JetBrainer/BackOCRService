package sqlstore

import (
	"database/sql"
)

type Store struct {
	db				*sql.DB
	userRepository	*UserRepository
}

// Constructor db
func New(db *sql.DB) *Store{
	return &Store{
		db: db,
	}
}