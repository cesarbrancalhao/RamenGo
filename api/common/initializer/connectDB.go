package initializer

import "api/database"

func ConnectDB() {
	database.ConnectDB()
}