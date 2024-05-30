package main

import (
	"api/cmd/common/initializer"
	"api/cmd/database"
	router "api/routers"
)

func init() {
	initializer.LoadEnv()
	database.ConnectDB()
}

func main () {
	router.StartRouter()
}