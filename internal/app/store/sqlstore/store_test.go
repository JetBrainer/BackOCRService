package sqlstore_test

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {

	if err := godotenv.Load(".env"); err != nil{
		log.Info().Msg("Unable to load env file")
	}
	databaseURL = os.Getenv("TESTDB")
	if databaseURL == ""{
		databaseURL = "host=localhost dbname=rest_test sslmode=disable"
	}

	os.Exit(m.Run())
}