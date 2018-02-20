package main

import (
	"log"

	"github.com/oherych/integration_testing_presentation/api"
)

func main() {

	cfg, err := api.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	api, err := api.NewAPi(cfg)
	if err != nil {
		log.Fatal(err)
	}

	api.DatabaseMigrate("db")

	if err := api.Run(); err != nil {
		log.Fatal(err)
	}
}
