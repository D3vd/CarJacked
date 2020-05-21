package casecnt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ca Controller) GetUnassignedCases(c *gin.Context) {
	cases, err := ca.M.GetAllUnassignedCases()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error while getting cases from DB",
		})
		return
	}

	c.JSON(http.StatusOK, cases)
	return
}

