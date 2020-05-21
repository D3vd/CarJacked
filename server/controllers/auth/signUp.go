package auth

import (
	"../../models"
	"../../models/requests"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (a Controller) SignUp(c *gin.Context) {

	// Create Empty Request body
	var req requests.SignUp

	// Parse Request Body
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong Request Body. Please try again",
		})
		return
	}

	//	Validate if all the request body fields are filled
	if req.Username == "" ||
		req.Password == "" ||
		req.Email == "" ||
		req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Please ensure that all the fields are filled",
		})
		return
	}

	// Create New User for Auth
	userID, err := a.M.CreateNewUser(req.Username, req.Password)

	if err != 0 {
		if err == 400 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Username is already present. Please choose another one",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server error while creating new user. Please try again",
			})
			return
		}
	}

	// Create new Officer Model
	officer := models.Officer{
		Name:     req.Name,
		Assigned: false,
	}

	if ok := a.M.CreateOfficer(officer, userID); ok != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error while creating a new Officer. Please try again",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"userID": userID,
	})
}
