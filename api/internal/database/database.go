package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    DB, err = gorm.Open(mysql.Open(os.Getenv("DB_CONN")), &gorm.Config{})
    if err!= nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}