package utility

import (
	"github.com/gin-gonic/gin"
)

// Controller ...
type Controller struct{}

// GetStatus used for health checks
func (ctrl Controller) GetStatus(c *gin.Context) {
	c.Status(200)
}
