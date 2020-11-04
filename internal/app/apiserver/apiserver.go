package apiserver

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Start server
func Start(config *Config) (*http.Server, *sql.DB){
	db, err := newDB(config.DBUrl)
	if err != nil{
		log.Error().Msg("Database URL ERROR")
		return nil, nil
	}
	r := newServer()
	serv := &http.Server{
		Addr: ":" + config.HttpPort,
		Handler: r,
	}
	go func() {
		if err := serv.ListenAndServe(); err != nil{
			log.Fatal().Err(err).Msg("Start up Failed")
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
