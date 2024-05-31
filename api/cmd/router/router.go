package router

import (
	routes "api/cmd/router/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()

	routes.OrderRoutes(router)
	routes.RecipeRoutes(router)
	routes.BrothRoutes(router)
	routes.ProteinRoutes(router)
	
	router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}