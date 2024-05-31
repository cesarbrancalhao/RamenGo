package router

import (
	"api/cmd/controller"
	middleware "api/pkg/middleware/cors"

	"github.com/gin-gonic/gin"
)
func RecipeRoutes(router *gin.Engine) {

	router.GET("/recipes", middleware.Cors, controller.GetAllRecipes)
	router.GET("/recipes/:id", middleware.Cors, controller.GetRecipe)
	router.POST("/recipes", middleware.Cors, controller.CreateRecipe)
	router.PUT("/recipes", middleware.Cors, controller.UpdateRecipe)
	router.DELETE("/recipes", middleware.Cors, controller.DeleteRecipe)
}