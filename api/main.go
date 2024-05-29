package main

import (
	"api/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main () {

	router.StartRouter()
}