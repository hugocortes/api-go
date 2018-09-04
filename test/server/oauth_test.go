package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hugocortes/paprika-api-go/modules"
	"github.com/stretchr/testify/assert"
)

func TestGetOAuthToken(t *testing.T) {
	testRouter := modules.InitRouter()

	data := url.Values{}
	req, err := http.NewRequest("POST", "/oauth/token", strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusBadRequest)
}
