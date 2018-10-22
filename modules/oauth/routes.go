package oauth

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes ...
func InitRoutes(router *gin.Engine) {
	controller := new(Controller)

	router.POST("/oauth/token", controller.GetOAuthToken)
}
