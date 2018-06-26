package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/api-go/modules"
	"github.com/hugocortes/api-go/plugins/keycloak"
	"github.com/joho/godotenv"
)

// init called before main() by golang runtime
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file %s", err)
	}
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Initialize singletons
	keycloak.Initialize()

	// API route initialization
	router := modules.InitRouter()
	router.Run(":" + os.Getenv("PORT"))
}
