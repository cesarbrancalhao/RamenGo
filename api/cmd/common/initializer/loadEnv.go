package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env.local")
    if err!= nil {
        log.Fatalf("Error loading.env file")
    }	
}