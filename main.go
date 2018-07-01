package main

import (
	"log"
	"os"

	"github.com/hugocortes/paprika-api/modules"
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
	// API route initialization
	router := modules.InitRouter()
	router.Run(":" + os.Getenv("PORT"))
}
