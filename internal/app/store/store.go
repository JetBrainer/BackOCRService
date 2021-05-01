package store

// Store user
type Store interface {
	User() UserRepository
}
