package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugocortes/paprika-api-go/data/sql"
)

// Controller for route users routes
type Controller struct {
	db sql.DB
}

// GetUserByID returns a user
func (ctrl Controller) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	user := ctrl.db.GetUserByID(userID)

	c.JSON(http.StatusOK, user)

	return
}
