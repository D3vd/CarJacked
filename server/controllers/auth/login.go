package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Login : Authenticate user using username and password
func (a Controller) Login(c *gin.Context) {
	session := sessions.Default(c)

	// Get username and password from the post body
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "Parameters can't be empty",
		})
		return
	}

	// TODO: Query DB with username for password

	// TODO: Check if passwords match
	if username != "admin" || password != "admin123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// TODO: Get userID and set it for session

	// Set and save session
	session.Set("userID", username)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"message": "Failed to save session",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "Successfully authenticated user",
	})
}
