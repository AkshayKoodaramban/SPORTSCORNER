package main

import (
	"fmt"
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
		fmt.Println("1")
	} else {
		fmt.Println("2")

		server.Start()
	}

}
