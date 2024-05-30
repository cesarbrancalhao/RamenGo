package router

import (
	controller "api/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()

	router.GET("/broths", controller.GetBroths)
	
	router.Run("localhost:8080")
}