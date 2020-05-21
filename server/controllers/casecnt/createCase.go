package casecnt

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"../../models"
	"../../models/requests"
)

// CreateCase : Create a new case and store to DB
func (ca Controller) CreateCase(c *gin.Context) {

	// Create empty request body
	var req requests.CaseRequest

	// Parse Request Body
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "Wrong Request Body. Please try again",
		})
		return
	}

	// Validate if all the required fields are present
	if req.User.Name == "" || req.User.Email == "" || req.User.Phone == "" || req.Car.Color == "" || req.Car.RegNo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "Please ensure that all the fields are filled",
		})
		return
	}

	// TODO: Check if the User's case is already present in the DB

	// Convert the request body to newCase Model

	newCase := models.Case{
		User: models.User{
			Name: req.User.Name,
			Email: req.User.Email,
			Phone: req.User.Phone,
		},
		Car: models.Car{
			Color: req.Car.Color,
			RegNo: req.Car.RegNo,
		},
		Active: true,
	}

	// TODO: Assign a Free Officer to the case

	// Write the onto the DB

	insertResult, err := ca.M.DB.Collection("case").InsertOne(context.Background(), newCase)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"message": "Error while inserting case to DB. Please try again.",
		})
	}

	log.Println(insertResult)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Successfully Added Case to DB",
	})

	return
}
