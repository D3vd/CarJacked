package casecnt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ca Controller) GetActiveCases(c *gin.Context) {

	cases, err := ca.M.GetAllActiveCases()

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
