package apiserver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Start server
func Start(config *Config) (*http.Server, *sql.DB){
	log.Info().Msg("Starting Database...")
	db, err := newDB(config.DBUrl)
	if err != nil{
		log.Error().Msg("Database URL ERROR")
		return nil, nil
	}
	r := newServer(config)
	serv := &http.Server{
		Addr: ":" + config.HttpPort,
		Handler: r,
	}
	log.Info().Msg("Starting server...")

	//go func() {
	//	if err := serv.ListenAndServe(); err == http.ErrServerClosed{
	//		log.Fatal().Err(err).Msg("Server closed")
	//	}
	//}()

	serv.ListenAndServe()
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
