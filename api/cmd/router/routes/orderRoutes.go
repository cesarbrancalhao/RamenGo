package router

import (
	controller "api/cmd/controller"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	
	router.GET("/orders", middleware.Cors, controller.GetOrderResponse)
	router.POST("/orders", middleware.Cors, controller.CreateOrderResponse)
	router.PUT("/orders", middleware.Cors, controller.UpdateOrderResponse)
	router.DELETE("/orders", middleware.Cors, controller.DeleteOrderResponse)
}