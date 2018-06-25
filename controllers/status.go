package controllers

import (
	"github.com/gin-gonic/gin"
)

// StatusController ...
type StatusController struct{}

// GetStatus used for health checks
func (ctrl StatusController) GetStatus(c *gin.Context) {
	c.Status(200)
}
