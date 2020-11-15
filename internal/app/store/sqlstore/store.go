package sqlstore

import (
	"database/sql"
	"github.com/JetBrainer/BackOCRService/internal/app/store"
)

type Store struct {
	Db             *sql.DB
	userRepository *UserRepository
}

// Constructor db
func New(db *sql.DB) *Store{
	return &Store{
		Db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil{
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
