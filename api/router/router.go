package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()
	router.GET("/broths", controller.GetBroth)

	router.Run("localhost:8080")
}