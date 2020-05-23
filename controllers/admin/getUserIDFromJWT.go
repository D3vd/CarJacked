package admin

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

// GetUserIDFromJWTBearer : Extract the JWT from the header
func (a Controller) GetUserIDFromJWTBearer (c *gin.Context) (string, error) {

	// Get token from header
	authHeader := c.GetHeader("Authorization")

	tokenSplit := strings.Split(authHeader, "Bearer")

	if len(tokenSplit) != 2 {
		return "", errors.New("bearer token not in proper format")
	}

	tokenString := strings.TrimSpace(tokenSplit[1])

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID, ok := claims["userID"].(string)
		if !ok {
			return "", err
		}

		return userID, nil
	}

	return "", err
}
