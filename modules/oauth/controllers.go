package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api/plugins/keycloak"
)

// Controller ...
type Controller struct{}

// GetOAuthToken used for different grant type authentication requests
func (ctrl Controller) GetOAuthToken(c *gin.Context) {
	var json TokenModel
	website := keycloak.GetWebInstance()

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch grantType := json.GrantType; grantType {
	case "password":
		transformedUMAToken, err := website.HandlePasswordGrant(json.Username, json.Password, json.Scope)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, transformedUMAToken)
		}
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
