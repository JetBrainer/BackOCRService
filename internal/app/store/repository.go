package store

type UserRepository interface {
	Create()
	Find()
	Update()
	Delete()
}