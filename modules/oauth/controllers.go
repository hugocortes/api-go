package oauth

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/api-go/plugins/keycloak"
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
		fmt.Println("password type")
		password := json.Password
		username := json.Username
		scope := json.Scope

		fmt.Println("Blah %s %s %s", username, password, scope)
		if scope == "" {
			fmt.Println("empty!")
		}

		data := url.Values{}
		data.Set("grant_type", "password")
		data.Add("username", username)
		data.Add("password", password)
		method := "POST"

		// this should be in struct
		keycloak := keycloak.GetWebInstance()
		data.Add("client_id", keycloak.ClientName)
		data.Add("client_secret", keycloak.ClientSecret)
		url := keycloak.UmaToken

		///
		///
		///
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("error: %s", err)
		}
		status := resp.StatusCode
		if status > 400 {
			fmt.Errorf("%v response when requesting user info", status)
		}
		///
		///
		///

		log.Println("*****")
		log.Println(status)
		log.Println("*****")

		break
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

	c.JSON(http.StatusOK, gin.H{"grant_type": json.GrantType})
	return
}
