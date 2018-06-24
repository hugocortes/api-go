package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t * testing.T) {
	router := gin.New()
	router.GET("/status", HealthCheckHandler)

	req, _ := http.NewRequest("GET", "/status", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}
