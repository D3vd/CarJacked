package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a Controller) CaseSolved(c *gin.Context) {

	userID, err := a.GetUserIDFromJWTBearer(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Bearer Token Format is Invalid",
		})
		return
	}

	// Mark case as not active
	err = a.M.MakeCaseNotActive(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Unable to update the case",
		})
		return
	}
	
	// Mark Officer as unassigned
	err = a.M.MakeOfficerUnassigned(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Unable to update the case",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Successfully updated case",
	})

}
