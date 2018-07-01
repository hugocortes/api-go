package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hugocortes/paprika-api/modules"
	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	testRouter := modules.InitRouter()

	req, err := http.NewRequest("GET", "/status", nil)

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}
