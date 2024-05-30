package router

import (
	controller "api/cmd/controller"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()
	
	// As for this is a very small project, I won't use a router file for each controller
	router.GET("/broths", middleware.Cors, controller.GetBroths)
	router.POST("/broths", middleware.Cors, controller.CreateBroth)
	
	router.Run("localhost:8080")
}
