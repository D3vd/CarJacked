package casecnt

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/R3l3ntl3ss/CarJacked/models"
)

// GetCaseWithID : Get One case filtered by ID
func (ca Controller) GetCaseWithID(c *gin.Context) {
	caseID := c.Param("caseID")

	// Check if caseID is present
	if caseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Please enter a valid Case ID",
		})
		return
	}

	var caseDoc models.Case

	//	Query DB for case
	errorCode := ca.M.GetCaseByID(caseID, &caseDoc)

	if errorCode != 0 {
		if errorCode == 400 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Unable to find case with id. Please check the ID and try again",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error while trying to case from DB. Please try again.",
		})
		return
	}

	c.JSON(http.StatusOK, caseDoc)

	return
}
