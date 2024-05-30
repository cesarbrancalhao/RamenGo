package main

import (
	"api/database"
	router "api/routers"
	"log"

	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err!= nil {
        log.Fatalf("Error loading.env file")
    }
}

func main () {
	database.InitDB()
	router.StartRouter()
}