package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
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

	// Database Close
	defer db.Close()

	// Server Shutdown
	defer serv.Shutdown(ctx)

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
