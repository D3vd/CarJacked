package casecnt

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"../../models"
)

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

	// Convert case ID string to primitive Object
	caseObjID, err := primitive.ObjectIDFromHex(caseID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Please enter a valid Case ID",
		})
		return
	}

	// Create filter for search
	filter := bson.M{
		"_id": caseObjID,
	}

	var caseDoc models.Case

	//	Query DB for case
	err = ca.M.DB.Collection("case").FindOne(context.TODO(), filter).Decode(&caseDoc)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
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
