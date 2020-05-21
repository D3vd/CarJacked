package server

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"../controllers/casecnt"
	"../libraries/mongo"
)

// NewRouter : Create the routes
func NewRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	m := &mongo.Mongo{}
	m.Init()

	router.GET("/test", func(c *gin.Context) {
		cur, err := m.DB.Collection("report").Find(context.Background(), bson.D{{}})

		if err != nil {
			log.Fatal(err)
		}

		var results []primitive.M

		for cur.Next(context.Background()) {
			var result bson.M
			e := cur.Decode(&result)
			if e != nil {
				log.Fatal(e)
			}

			results = append(results, result)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.Background())

		log.Println(results)

		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	caseRouter := router.Group("/case")
	{
		c := casecnt.Controller{
			M: m,
		}

		caseRouter.POST("", c.CreateCase)
		caseRouter.GET("/:caseID", c.GetCaseWithID)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
