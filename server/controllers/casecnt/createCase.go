package casecnt

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"../../models/requests"
)

// CreateCase : Create a new case and store to DB
func (ca Controller) CreateCase(c *gin.Context) {

	// Create empty request body
	var req requests.CaseRequest

	// Parse Request Body
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Request Body. Please try again",
		})
	}

	// TODO: Validate if all the required fields are present

	// TODO: Convert the request body to Model

	// TODO: Write the onto the DB

	log.Println(req)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}
