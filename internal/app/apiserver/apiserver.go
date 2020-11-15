package apiserver

import (
	"database/sql"
	"github.com/JetBrainer/BackOCRService/internal/app/store/sqlstore"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

// Start server
func Start(config *Config) (*http.Server, *sql.DB){
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

func newDB(databaseURL string) (*sql.DB,error){

	db, err := sql.Open("postgres",databaseURL)
	if err != nil{
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}
