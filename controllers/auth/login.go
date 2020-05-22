package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"../../models/requests"
)

// Login : Authenticate user using username and password
func (a Controller) Login(c *gin.Context) {
	session := sessions.Default(c)

	// Create Empty Request body
	var req requests.Login

	// Parse Request Body
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Wrong Request Body. Please try again",
		})
		return
	}

	username := req.Username
	password := req.Password

	// Check if the user is already logged in
	sessionID := session.Get("userID")

	if sessionID != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "User is already logged in",
		})
		return
	}

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Parameters can't be empty",
		})
		return
	}

	// Query DB with username for password
	userPassword, userID, err := a.M.GetUserPasswordAndID(username)

	// Error Handling for DB username Error
	if err != 0 {
		if err == 400 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Username is not present in the DB. Please create a new account",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Error while querying DB for username. Please try again",
			})
			return
		}
	}

	// Check if passwords match
	if password != userPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Authentication Failed. Wrong password",
		})
		return
	}

	// Get userID and set it for session
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to save session",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully authenticated user",
	})
}
