package modules

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api-go/data/sql"
	"github.com/hugocortes/paprika-api-go/modules/oauth"
	"github.com/hugocortes/paprika-api-go/modules/users"
	"github.com/hugocortes/paprika-api-go/modules/utility"
)

// InitRouter ..
func InitRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()

	utility.InitUtilityRoutes(router)
	oauth.InitOAuthRoutes(router)
	users.InitUsersRoutes(router, db)

	router.NoRoute(notFoundHandler)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Authorization", "Accept", "x-access-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	return router
}

func notFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"status": 404, "error": "Not Found", "message": "Not Found"})
}
