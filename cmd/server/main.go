package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Toml config path file
var configPath string

// Parsing values
func init(){
	flag.StringVar(&configPath,"config","config/apiserver.toml","path to config file")
}

func main(){
	// Parse flags
	flag.Parse()

	// Configs
	config := apiserver.InitConfig()

	// Decode Toml file and record
	_, err := toml.DecodeFile(configPath,&config)
	if err != nil{
		log.Fatal().Err(err).Msg("Decode file Failed")
	}

	// context to shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Server
	serv, db := apiserver.Start(config)


	// Server Shutdown
	defer func(serv *http.Server) {
		if err := serv.Shutdown(ctx); err != nil{
			log.Info().Msg("Server Shutdown error")
		}
	}(serv)

	// Database Close
	defer func(db *mongo.Client) {
		if err := db.Disconnect(context.Background()); err != nil{
			log.Info().Msg("Error db closing")
		}
	}(db)

	// Signal
	handleSignals()
}

// Graceful Shutdown
func handleSignals() os.Signal{
	sigChan := make(chan os.Signal,1)
	signal.Notify(sigChan,os.Interrupt, syscall.SIGTERM)
	for {
		sig := <-sigChan
		switch sig {
		case os.Interrupt:
			log.Info().Msg("Graceful interrupt server")
			return sig
		case syscall.SIGTERM:
			log.Info().Msg("Graceful Kill server")
			return sig
		}
	}
}
