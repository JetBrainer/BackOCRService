package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/JetBrainer/BackOCRService/internal/app/apiserver"
	"log"
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
		log.Fatal(err)
	}

	// context to shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Server
	serv, db := apiserver.Start(config)

	defer func() {
		if err := db.Close(); err != nil{
			log.Fatal(err)
			return
		}
	}()

	defer func() {
		if err := serv.Shutdown(ctx); err != nil{
			log.Fatal(err)
		}
	}()

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
			log.Println("Graceful Interrupt Shutdown")
			return sig
		case syscall.SIGTERM:
			log.Println("Graceful Kill Shutdown")
			return sig
		}
	}
}
