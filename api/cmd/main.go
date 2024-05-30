package main

import (
	router "api/cmd/routers"
	"api/common/initializer"
	"api/internal/database"
)

func init() {
	initializer.LoadEnv()
	database.ConnectDB()
}

func main () {
	router.StartRouter()
}