package oauth

import (
	"github.com/gin-gonic/gin"
)

// InitOAuthRoutes ...
func InitOAuthRoutes(router *gin.Engine) {
	controller := new(Controller)

	router.POST("/oauth/token", controller.GetOAuthToken)
}
