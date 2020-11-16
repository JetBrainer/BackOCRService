package store

import "github.com/JetBrainer/BackOCRService/internal/app/model"

type UserRepository interface {
	Create(*model.User) 	error
	FindByEmail(string) 	(*model.User,error)
	Find(int)				(*model.User,error)
	UpdateUser(*model.User) error
	DeleteUser(string) 		error
	CheckToken(string)		error
}