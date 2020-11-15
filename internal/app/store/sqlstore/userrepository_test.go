package sqlstore_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JetBrainer/BackOCRService/internal/app/model"
	"github.com/JetBrainer/BackOCRService/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock := sqlstore.MockDB(t)

	user := model.TestUser(t)
	rdbms := &sqlstore.Store{Db: db}
	defer func() {
		rdbms.Db.Close()
	}()

	if err := user.Validate(); err != nil{
		t.Fail()
	}
	if err := user.BeforeCreate(); err != nil{
		t.Fail()
	}

	query := "INSERT INTO acc \\(email,encpassword,organization,token\\) VALUES \\(\\$1,\\$2,\\$3,\\$4\\) RETURNING id"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(user.Email,user.EncryptedPassword,user.Organization,user.Token).
		WillReturnResult(sqlmock.NewResult(1,1))

	assert.NoError(t,rdbms.User().Create(user))
}

func TestUserRepository_FindByEmail(t *testing.T) {

}

func TestUserRepository_Find(t *testing.T) {

}

func TestUserRepository_UpdateUser(t *testing.T) {

}

func TestUserRepository_DeleteHandler(t *testing.T) {

}