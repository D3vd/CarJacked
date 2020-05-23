package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DEPRECATED : This function uses logs out the user for Session based Authentication
//              But now the auth method is changed to JWT
// SessionLogout : Logout the user by removing the userID from session
func (a Controller) SessionLogout(c *gin.Context) {
	session := sessions.Default(c)

	// Get userID from session
	userID := session.Get("userID")

	// Check if UserID is present
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "User is not logged in",
		})
		return
	}

	// Check the UserID from session and save it
	session.Delete("userID")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error while deleting userID",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully logged out",
	})
}
