package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api-go/modules"
	"github.com/hugocortes/paprika-api-go/modules/utility"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var utilityApp struct {
	router *gin.Engine
}

func utilitySetup() {
	if utilityApp.router != nil {
		return
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("did not load from .env")
	}

	utilityApp.router = modules.GetDefaultRouter()
	utility.InitRoutes(utilityApp.router)

	return
}

func TestGetStatus(t *testing.T) {
	utilitySetup()

	req, err := http.NewRequest("GET", "/status", nil)

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	utilityApp.router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}
