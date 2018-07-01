/**
 * Initializes singleton Keycloak
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
	clientID     string
	clientName   string
	clientSecret string
	clientToken  string

	// Admin Routes
	adminUsers     string
	adminResources string
	adminEvaluate  string

	// UMA Routes
	umaAuth     string
	umaToken    string
	umaUserInfo string
	umaResource string
}

var websiteInstance *keycloak
var once sync.Once

// GetWebInstance returns a singleton keycloak struct
func GetWebInstance() *keycloak {
	if websiteInstance == nil {
		initialize()
	}
	return websiteInstance
}

// Initialize keycloak clients as singletons
func initialize() {
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
		clientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		clientName:   os.Getenv("KEYCLOAK_CLIENT_NAME"),
		clientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}

func (client *keycloak) setUmaTokenPath() {
	client.umaToken = client.baseURL + "/auth/realms/" + client.realm + "/protocol/openid-connect/token"
}
