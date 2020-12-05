package sqlstore

import (
	"context"
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User Repo for db work
type UserRepository struct {
	store *Store
}

// Create User
func (r *UserRepository) Create(u *model.User) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
	collect := r.store.Db.Database("testing").Collection("acc")
	res, err := collect.InsertOne(ctx, bson.M{
		"email":				u.Email,
		"encryptedpassword":	u.EncryptedPassword,
		"organization":			u.Organization,
		"token":				u.Token,
	})

	u.ID = res.InsertedID.(primitive.ObjectID)
	return err
}

// Find By Email
func (r *UserRepository) FindByEmail(email string)(*model.User,error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	u := &model.User{}
	collect := r.store.Db.Database("testing").Collection("acc")

	err := collect.FindOne(ctx,bson.M{
		"email":email,
	}).Decode(&u)
	return u, err
}

// Find by id
func (r *UserRepository) Find(id primitive.ObjectID) (*model.User,error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	u := &model.User{}
	collect := r.store.Db.Database("testing").Collection("acc")

	err := collect.FindOne(ctx,bson.M{
		"_id":id,
	}).Decode(&u)
	return u, err
}

// Updated User
func (r *UserRepository) UpdateUser(u *model.User) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := u.Validate(); err != nil{
		return err
	}
	if err := u.BeforeCreate(); err != nil{
		return err
	}

	collect := r.store.Db.Database("testing").Collection("acc")

	filter := bson.M{
		"email":u.Email,
	}
	update := bson.M{
		"$set": bson.M{
			"encryptedpassword":u.EncryptedPassword,
		},
	}
	err := collect.FindOneAndUpdate(ctx,filter,update).Decode(&u)
	return err
}

// Delete User
func (r *UserRepository) DeleteUser(email string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collect := r.store.Db.Database("testing").Collection("acc")

	_, err := collect.DeleteOne(ctx, bson.M{
		"email":email,
	})
	if err != nil{
		return err
	}
	log.Info().Msg("Document Changed")
	return nil
}

// Check Token
func (r *UserRepository) CheckToken(Token string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	u := &model.User{}
	collect := r.store.Db.Database("testing").Collection("acc")

	err := collect.FindOne(ctx,bson.M{
		"email":Token,
	}).Decode(&u)
	return err
}