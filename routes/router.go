package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	versionnedRouter1 := router.Group("/v1")

	addOrderRoutes(versionnedRouter1)

	return router
}
