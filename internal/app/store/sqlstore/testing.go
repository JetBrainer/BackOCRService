package sqlstore

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog/log"
	"testing"
)

// Sql mock
func MockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock){
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil{
		log.Fatal().Msg("Sql Mock Error")
	}
	return db, mock
}