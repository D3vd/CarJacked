package admin

import (
	"github.com/R3l3ntl3ss/CarJacked/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a Controller) CheckForNewCase(c *gin.Context) {

	userID, err := a.GetUserIDFromJWTBearer(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Bearer Token Format is Invalid",
		})
		return
	}

	var assignedCase models.Case

	ok := a.M.GetCaseByOfficerID(userID, &assignedCase)

	if ok == 500 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Error while getting officer cases",
		})
		return
	}

	if ok == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"case": assignedCase,
		})
		return
	}

	assigned, newCase, err := a.M.AssignOfficerToCase(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Error while Assigning Officer to new case",
		})
		return
	}

	if assigned == false {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"case": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"case": newCase,
	})
}
