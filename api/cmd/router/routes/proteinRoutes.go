package router

import (
	"api/cmd/controller"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)
func ProteinRoutes(router *gin.Engine) {

	router.GET("/proteins", middleware.Cors, controller.GetProteins)
	router.POST("/proteins", middleware.Cors, controller.CreateProtein)
	router.PUT("/proteins", middleware.Cors, controller.UpdateProtein)
	router.DELETE("/proteins", middleware.Cors, controller.DeleteProtein)
}