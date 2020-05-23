package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/R3l3ntl3ss/CarJacked/models/requests"

	jwtLib "github.com/dgrijalva/jwt-go"
)

// Login : Authenticate user using username and password
func (a Controller) Login(c *gin.Context) {

	secret := os.Getenv("SECRET")

	// Create the token
	token := jwtLib.New(jwtLib.GetSigningMethod("HS256"))

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

	// Set some claims
	token.Claims = jwtLib.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, ok := token.SignedString([]byte(secret))

	if ok != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Could not generate token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully authenticated user",
		"token":   tokenString,
	})
}
