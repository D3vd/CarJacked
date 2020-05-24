package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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

	// Convert userID string to primitive Object ID
	officerObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": "Internal server Error while converting userID to object",
		})
		return
	}

	unassignedCases, err := a.M.GetAllUnassignedCases()

	log.Println(unassignedCases)

	if unassignedCases != nil {
		if len(unassignedCases) > 0 {

			newCase := unassignedCases[0]

			// Assign case to officer
			err = a.M.UpdateCaseWithOfficerID(newCase.ID, officerObjID)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"message": "Internal server Error while converting userID to object",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "Successfully updated case",
				"case":    newCase,
			})
			return
		}
	}

	// Mark Officer as unassigned
	err = a.M.MakeOfficerUnassigned(userID)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Unable to update the case",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Successfully updated case",
		"case": nil,
	})

}
