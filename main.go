package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/status", status)
	r.Run()
}

func status(c *gin.Context) {
	c.Status(200)
}
