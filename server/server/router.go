package server

import (
	"../controllers/casecnt"
	"../libraries/mongo"
	"github.com/gin-gonic/gin"
)

// NewRouter : Create the routes
func NewRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	m := &mongo.Mongo{}
	m.Init()

	caseRouter := router.Group("/case")
	{
		c := casecnt.Controller{
			M: m,
		}

		caseRouter.POST("", c.CreateCase)
		caseRouter.GET("/active", c.GetActiveCases )
		caseRouter.GET("/unassigned", c.GetActiveCases )
		caseRouter.GET("/id/:caseID", c.GetCaseWithID)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
