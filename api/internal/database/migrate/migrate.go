package main

import (
	"api/common/initializer"
	"api/internal/database"
	"api/internal/models"
)

func init() {
	initializer.LoadEnv()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.Broth{})
	database.DB.AutoMigrate(&models.Protein{})
	database.DB.AutoMigrate(&models.Recipe{})
	database.DB.AutoMigrate(&models.OrderResponse{})
}