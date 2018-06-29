package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl Controller) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	return
}
