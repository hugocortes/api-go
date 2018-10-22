package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api-go/modules"
	"github.com/hugocortes/paprika-api-go/modules/oauth"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var oauthApp struct {
	router *gin.Engine
}

func oauthSetup() {
	if oauthApp.router != nil {
		return
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("did not load from .env")
	}

	// api
	oauthApp.router = modules.GetDefaultRouter()
	oauth.InitRoutes(oauthApp.router)

	return
}

func TestGetOAuthToken(t *testing.T) {
	oauthSetup()

	data := url.Values{}
	req, err := http.NewRequest("POST", "/oauth/token", strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	oauthApp.router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusBadRequest)
}
