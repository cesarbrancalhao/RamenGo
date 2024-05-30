package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))

    if err!= nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Successfully connected to the database")
}