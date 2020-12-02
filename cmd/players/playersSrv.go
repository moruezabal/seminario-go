package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moruezabal/seminario-go/internal/config"
	"github.com/moruezabal/seminario-go/internal/database"
	"github.com/moruezabal/seminario-go/internal/service/players"
)

func main() {

	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := players.New(db, cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg

}
