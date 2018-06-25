package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocortes/api-go/controllers"
)

func initStatusRoutes(router *gin.Engine) {
	status := new(controllers.StatusController)
	router.GET("/status", status.GetStatus)
}
