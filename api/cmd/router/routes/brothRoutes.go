package router

import (
	"api/cmd/controller"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)

func BrothRoutes(router *gin.Engine) {
	
    router.GET("/broths", middleware.Cors, controller.GetBroths)
    router.POST("/broths", middleware.Cors, controller.CreateBroth)
    router.PUT("/broths", middleware.Cors, controller.UpdateBroth)
    router.DELETE("/broths", middleware.Cors, controller.DeleteBroth)
}