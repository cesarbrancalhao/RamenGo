package router

import (
	controller "api/cmd/controller"
	"api/pkg/handler"
	cors "api/pkg/middleware/cors"
	httpHandler "api/pkg/middleware/http"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	
	router.GET("/orders", cors.Cors, controller.GetAllOrderResponses)
	router.GET("/orders/:id", cors.Cors, controller.GetOrderResponse)
	router.POST("/orders",
		cors.Cors,
		handler.HandleOrderRequest,
		httpHandler.GenerateId,
		controller.CreateOrder)
	router.PUT("/orders", cors.Cors, controller.UpdateOrderResponse)
	router.DELETE("/orders", cors.Cors, controller.DeleteOrderResponse)
}