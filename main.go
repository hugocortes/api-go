package main

import (
	"fmt"
	"os"

	"github.com/hugocortes/paprika-api-go/data/sql"
	"github.com/hugocortes/paprika-api-go/modules"
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
	// DB connection
	db := sql.NewConnection(os.Getenv("MYSQL_DSN"), os.Getenv("GIN_MODE") == "debug")

	// API route initialization
	router := modules.InitRouter(db)
	router.Run(":" + os.Getenv("PORT"))
}
