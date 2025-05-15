package main

import (
	"log"
	"scarlet/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ENV] -> Error loading .env file. Error: \n", err.Error())
	}

	api.Start()
}
