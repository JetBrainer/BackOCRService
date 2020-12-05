package apiserver

import (
	"context"
	"github.com/JetBrainer/BackOCRService/internal/app/store/sqlstore"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

// Start server
func Start(config *Config) (*http.Server, *mongo.Client){
	log.Info().Msg("Starting Database...")
	db, err := newDB(config.DBUrl)
	if err != nil{
		log.Error().Msg("Database URL ERROR")
		return nil, nil
	}
	store := sqlstore.New(db)
	r := newServer(store,config)
	serv := &http.Server{
		Addr: ":" + config.HttpPort,
		Handler: r,
		ReadTimeout: 7 *time.Second,
		WriteTimeout: 10*time.Second,
	}
	log.Info().Msg("Starting server...")

	go func() {
		if err := serv.ListenAndServe(); err == http.ErrServerClosed{
			log.Fatal().Err(err).Msg("Server closed")
		}
	}()


	return serv, db
}

func newDB(databaseURL string) (*mongo.Client,error){

	db, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil{
		return nil, err
	}
	if err := db.Ping(context.Background(),readpref.Primary()); err != nil{
		return nil, err
	}
	return db, nil
}
