package sqlstore

import (
	"database/sql"
	"errors"
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/rs/zerolog/log"
)

// User Repo for db work
type UserRepository struct {
	store *Store
}

// Create User
func (r *UserRepository) Create(u *model.User) error{
	// Validation of User
	if err := u.Validate(); err != nil{
		log.Info().Msg("Validation ERROR")
		return err
	}

	// Encryption Password
	if err := u.BeforeCreate(); err != nil{
		log.Info().Msg("Before create ERROR")
		return err
	}

	// Generate Token for user
	u.Token = model.TokenGenerator()

	return r.store.Db.QueryRow(
		"INSERT INTO acc(email,encpassword,organization,token) VALUES ($1,$2,$3,$4) RETURNING id",
		u.Email,u.EncryptedPassword,u.Organization,u.Token).Scan(&u.ID)
}

// Find By Email
func (r *UserRepository) FindByEmail(email string)(*model.User,error){
	u := &model.User{}

	if err := r.store.Db.QueryRow(
		"SELECT id,email,organization,token FROM acc WHERE email=$1",email).
		Scan(&u.ID,&u.Email,&u.Organization,&u.Token);
	err != nil{
		if err == sql.ErrNoRows{
			return nil, errors.New("SQL NO ROWS")
		}
		return nil, err
	}

	return u, nil
}

// Find by id
func (r *UserRepository) Find(id int) (*model.User,error){
	u := &model.User{}
	if err := r.store.Db.QueryRow(
		"SELECT id,email,encpassword,organization,token FROM acc WHERE id=$1",id).
		Scan(&u.ID,&u.Email,&u.Password,&u.Organization,&u.Token);
		err != nil{
			if err == sql.ErrNoRows{
				return nil, errors.New("SQL NO ROWS")
			}
		return nil, err
	}

	return u, nil
}

// Updated User
func (r *UserRepository) UpdateUser(u *model.User) error{
	if err := u.Validate(); err != nil{
		return err
	}
	if err := u.BeforeCreate(); err != nil{
		return err
	}

	return r.store.Db.QueryRow(
		"UPDATE acc SET encpassword=$1 WHERE email=$2 RETURNING id",u.EncryptedPassword,u.Email).
		Scan(&u.ID)
}

// Delete User
func (r *UserRepository) DeleteUser(email string) error{
	_, err := r.store.Db.Exec("DELETE FROM acc WHERE email = $1",email)
	if err != nil{
		return err
	}
	return nil
}

// Check Token
func (r *UserRepository) CheckToken(Token string) error{
	if _,err := r.store.Db.Exec(
		"SELECT token FROM acc WHERE token=$1", Token); err != nil{
		if err == sql.ErrNoRows{
			return errors.New("SQL NO ROWS")
		}
		return err
	}
	return nil
}