package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"github.com/rs/zerolog/log"
)

func main() {
	// Parse flags
	flag.Parse()

	// Configs
	config, err := apiserver.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Decode file Failed")
	}
	// context to shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Server
	serv, db := apiserver.Start(config)

	// Server Shutdown
	defer serv.Shutdown(ctx)

	// Database Close
	defer db.Close()

	// Signal
	handleSignals()
}

// Graceful Shutdown
func handleSignals() os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
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
