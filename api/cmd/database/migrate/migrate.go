package main

import (
	"api/cmd/common/initializer"
	"api/cmd/database"
	"api/internal/models"
)

func init() {
	initializer.LoadEnv()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.Broth{})
	database.DB.AutoMigrate(&models.Protein{})
	database.DB.AutoMigrate(&models.OrderResponse{})
}