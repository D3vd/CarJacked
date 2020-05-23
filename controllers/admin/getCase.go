package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCase : Get the case of the officer whose userID is in the JWT
func (a Controller) GetCase(c *gin.Context) {

	userID, err := a.GetUserIDFromJWTBearer(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"message": "Bearer Token Format is Invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"userID": userID,
	})

}
