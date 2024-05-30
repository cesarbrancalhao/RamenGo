package main

import (
	"api/common/initializer"
	router "api/routers"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}
    

func main () {
	router.StartRouter()
}