package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// Registrated user
type User struct {
	ID					int
	Email				string
	Organization		string
	Password 			string
	EncryptedPassword	string
	Token				string
}


// Validation User
func (u *User) Validate() error{
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(8,100),
			validation.Match(regexp.MustCompile("^[A-Z]"))))
}

//
func (u *User) BeforeCreate() error{
	if len(u.Password) > 0{
		enc, err := encryptString(u.Password)
		if err != nil{
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error){
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil{
		return "", err
	}

	return string(b), nil
}