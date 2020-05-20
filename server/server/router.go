package server

import (
	"github.com/gin-gonic/gin"
)

// NewRouter : Create the routes
func NewRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
