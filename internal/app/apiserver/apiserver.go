package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"

	_ "github.com/lib/pq"
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
	go func() {
		if err := serv.ListenAndServe(); err != nil{
			log.Fatal().Err(err).Msg("Start server Failed")
		}
	}()
	return serv, db
}

func newDB(databaseURL string) (*sql.DB,error){
	fmt.Println(databaseURL)
	db, err := sql.Open("postgres",databaseURL)
	if err != nil{
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}
