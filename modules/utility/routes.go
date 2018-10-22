package utility

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes ...
func InitRoutes(router *gin.Engine) {
	controller := new(Controller)

	router.GET("/status", controller.GetStatus)
}
