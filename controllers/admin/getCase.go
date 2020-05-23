package admin

import (
	"github.com/R3l3ntl3ss/CarJacked/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCase : Get the case of the officer whose userID is in the JWT
func (a Controller) GetCase(c *gin.Context) {

	userID, err := a.GetUserIDFromJWTBearer(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Bearer Token Format is Invalid",
		})
		return
	}

	var assignedCase models.Case

	errCode := a.M.GetCaseByOfficerID(userID, &assignedCase)

	if errCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": "Internal Server Error while getting case. PLease try again",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"case": assignedCase,
	})

}
