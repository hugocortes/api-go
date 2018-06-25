package routes

import (
	"github.com/gin-gonic/gin"
)

// InitRouter defines all routes
func InitRouter() *gin.Engine {
	router := gin.Default()

	initStatusRoutes(router)
	router.NoRoute(notFoundHandler)

	return router
}

func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"status": 404, "error": "Not Found", "message": "Not Found"})
}
