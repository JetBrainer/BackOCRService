package sqlstore_test

import (
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/JetBrainer/BackOCRService/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("acc")
	s := sqlstore.New(db)

	assert.NoError(t,s.User().Create(model.TestUser(t)))
	assert.NotNil(t, model.TestUser(t))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("acc")
	s := sqlstore.New(db)
	u := model.TestUser(t)

	email := "user@example.org"
	assert.NoError(t,s.User().Create(u))
	assert.NotNil(t, u)

	_, err := s.User().FindByEmail(email)
	assert.NoError(t,err)
	assert.NotNil(t,u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("acc")
	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t,s.User().Create(u))
	assert.NotNil(t, u)

	_, err := s.User().Find(u.ID)
	assert.NoError(t,err)
	assert.NotNil(t,u)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("acc")
	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t,s.User().Create(u))
	assert.NotNil(t, u)

	assert.NoError(t,s.User().UpdateUser(u))

}

func TestUserRepository_DeleteHandler(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("acc")
	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t,s.User().Create(u))
	assert.NotNil(t, u)

	assert.NoError(t,s.User().DeleteUser(u.Email))
}

func TestUserRepository_CheckToken(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("acc")
	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t,s.User().Create(u))
	assert.NotNil(t, u)

	assert.NoError(t,s.User().CheckToken(u.Token))
}