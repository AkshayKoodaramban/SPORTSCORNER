package main

import (
	"log"
	"sportscorner/pkg/config"
	"sportscorner/pkg/di"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		err := server.Start()
		if err != nil {
			log.Fatal("problem starting server")
		}
	}

}
