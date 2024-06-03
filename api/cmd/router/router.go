package router

import (
	routes "api/cmd/router/routes"
	middleware "api/pkg/middleware/cors"
	"os"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()

    router.Use(middleware.Cors)

	routes.OrderRoutes(router)
	routes.RecipeRoutes(router)
	routes.BrothRoutes(router)
	routes.ProteinRoutes(router)
	
	router.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}