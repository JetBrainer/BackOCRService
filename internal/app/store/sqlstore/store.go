package sqlstore

import (
	"github.com/JetBrainer/BackOCRService/internal/app/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	Db             *mongo.Client
	userRepository *UserRepository
}

// Constructor db
func New(db *mongo.Client) *Store{
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
