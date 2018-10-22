package users

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api-go/data/sql"
)

// InitUsersRoutes ...
func InitUsersRoutes(router *gin.Engine, db *sql.DB) {
	controller := Controller{
		db: *db,
	}

	router.GET("/users/:id", controller.GetUserByID)
}
