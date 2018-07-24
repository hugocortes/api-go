package main

import (
	"fmt"
	"os"

	"github.com/hugocortes/paprika-api/modules"
	"github.com/joho/godotenv"
)

// init called before main() by golang runtime
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("did not load from .env")
	}
}

func main() {
	// API route initialization
	router := modules.InitRouter()
	router.Run(":" + os.Getenv("PORT"))
}
