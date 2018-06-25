package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter defines all routes
func InitRouter() *gin.Engine {
	router := gin.Default()

	initStatusRoutes(router)

	router.NoRoute(notFoundHandler)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Authorization", "Accept", "x-access-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}

func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"status": 404, "error": "Not Found", "message": "Not Found"})
}
