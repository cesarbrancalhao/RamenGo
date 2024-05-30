package router

import (
	controller "api/cmd/controller"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()

	router.GET("/broths", middleware.Cors, controller.GetBroths)
	
	router.Run("localhost:8080")
}
