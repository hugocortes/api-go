/**
 * Singleton Keycloak
 * Implemented with thread safety
 */

package keycloak

import (
	"os"
	"sync"
)

// Keycloak ...
type keycloak struct {
	// Configuration
	baseURL      string
	realm        string
	ClientID     string
	ClientName   string
	ClientSecret string
	clientToken  string

	// Admin Routes
	AdminUsers     string
	AdminResources string
	AdminEvaluate  string

	// UMA Routes
	UmaAuth     string
	UmaToken    string
	UmaUserInfo string
	UmaResource string
}

var websiteInstance *keycloak
var once sync.Once

// Initialize keycloak clients as singletons
func Initialize() {
	once.Do(func() {
		websiteInstance = initializeWebsite()

		// Go through setters for realm
		websiteInstance.setUmaTokenPath()
	})
}

func initializeWebsite() *keycloak {
	return &keycloak{
		// base configuration
		baseURL:      os.Getenv("KEYCLOAK_BASE"),
		realm:        os.Getenv("KEYCLOAK_REALM"),
		ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientName:   os.Getenv("KEYCLOAK_CLIENT_NAME"),
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}

// GetWebInstance returns a singleton keycloak struct
func GetWebInstance() *keycloak {
	return websiteInstance
}

func (client *keycloak) setUmaTokenPath() {
	client.UmaToken = client.baseURL + "/auth/realms/" + client.realm + "/protocol/openid-connect/token"
}
