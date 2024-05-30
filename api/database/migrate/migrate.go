package main

import (
	"api/common/initializer"
	"api/database"
	"api/models"
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