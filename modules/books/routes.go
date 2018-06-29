package books

import "github.com/gin-gonic/gin"

// InitBookRoutes ...
func InitBookRoutes(router *gin.Engine) {
	controller := new(Controller)

	router.GET("/books", controller.GetBooks)
}
