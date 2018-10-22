package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller ...
type Controller struct{}

// GetOAuthToken used for different grant type authentication requests
func (ctrl Controller) GetOAuthToken(c *gin.Context) {
	var json TokenModel

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch grantType := json.GrantType; grantType {
	case "password":
		fmt.Println("password grant type")
		return
	case "client_credentials":
		fmt.Println("client cred type")
		break
	case "authorization_code":
		fmt.Println("auth code")
		break
	case "impersonation":
		fmt.Println("impersonate type")
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid grant type"})
		return
	}

	return
}
