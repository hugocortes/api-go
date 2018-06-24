package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/status", HealthCheckHandler)
	r.Run()
}

func HealthCheckHandler(c *gin.Context) {
	c.Status(200)
}
