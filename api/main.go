package main

import (
	"api/common/initializer"
	"api/database"
	router "api/routers"
)

func init() {
	initializer.LoadEnv()
	database.ConnectDB()
}

func main () {
	router.StartRouter()
}