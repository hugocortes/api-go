package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api-go/api"
	"github.com/hugocortes/paprika-api-go/api/users"
	"github.com/hugocortes/paprika-api-go/data/sql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var usersApp struct {
	db     *sql.DB
	router *gin.Engine
}

func usersSetup() {
	if usersApp.router != nil {
		return
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("did not load from .env")
	}

	// DB connection
	usersApp.db = sql.NewConnection(os.Getenv("MYSQL_DSN"), true)
	usersApp.router = api.GetDefaultRouter()
	users.InitRoutes(usersApp.router, usersApp.db)

	return
}

func TestGetUserById(t *testing.T) {
	usersSetup()

	userID := "95c8f54e-4bb4-4f6b-9cb3-2f198bfbe07c"

	req, err := http.NewRequest("GET", "/users/"+userID, nil)

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	usersApp.router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
}
