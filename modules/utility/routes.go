package utility

import (
	"github.com/gin-gonic/gin"
)

// InitUtilityRoutes ...
func InitUtilityRoutes(router *gin.Engine) {
	controller := new(Controller)

	router.GET("/status", controller.GetStatus)
}
