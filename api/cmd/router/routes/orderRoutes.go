package router

import (
	controller "api/cmd/controller"
	"api/pkg/handler"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	
	router.GET("/orders", middleware.Cors, controller.GetAllOrderResponses)
	router.GET("/orders/:id", middleware.Cors, controller.GetOrderResponse)
	router.POST("/orders", middleware.Cors, controller.CreateOrderResponse)
	router.POST("/order",
		middleware.Cors,
		handler.HandleOrderRequest,
		controller.GetOrderResponseByRequest)
	router.PUT("/orders", middleware.Cors, controller.UpdateOrderResponse)
	router.DELETE("/orders", middleware.Cors, controller.DeleteOrderResponse)
}