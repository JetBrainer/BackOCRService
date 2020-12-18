package sqlstore

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"

	_ "github.com/lib/pq"
)

// Sql mock
func TestDB(t *testing.T, databaseURL string) (*mongo.Client, func(...string)){
	t.Helper()

	db, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseURL))
	if err != nil{
		t.Fatal(err)
	}
	collection := db.Database("test_db").Collection("diploma")
	if err := db.Ping(context.TODO(), readpref.Primary()); err != nil{
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables)>0{
			if err := collection.Drop(context.TODO()); err != nil{
				log.Info().Err(err).Msg("Execution Error")
			}
		}

		db.Disconnect(context.TODO())
	}
}