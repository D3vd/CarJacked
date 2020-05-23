package server

import (
	"github.com/R3l3ntl3ss/CarJacked/controllers/admin"
	"github.com/R3l3ntl3ss/CarJacked/controllers/auth"
	"github.com/R3l3ntl3ss/CarJacked/controllers/casecnt"
	"github.com/gin-gonic/contrib/jwt"
	"os"

	"github.com/R3l3ntl3ss/CarJacked/libraries/mongo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// NewRouter : Create the routes
func NewRouter() *gin.Engine {

	router := gin.New()

	// Add Sessions Middleware
	router.Use(sessions.Sessions("mySession", sessions.NewCookieStore([]byte("secret"))))

	// Add Logging and Recovery Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	// Get JWT Auth Secret
	secret := os.Getenv("SECRET")

	// Initialize MongoDB server
	m := &mongo.Mongo{}
	m.Init()

	// Routes for Authentication
	authRouter := router.Group("/auth")
	{
		a := auth.Controller{
			M: m,
		}

		authRouter.POST("/login", a.Login)
		authRouter.GET("/logout", a.Logout)
		authRouter.POST("/signUp", a.SignUp)
	}

	// Routes for admin operations [Requires Authentication]
	adminRouter := router.Group("/admin")
	adminRouter.Use(jwt.Auth(secret))
	{
		a := admin.Controller{
			M: m,
		}

		adminRouter.GET("/test", a.Test)
	}

	// Routes for creating and seeing cases [Public]
	caseRouter := router.Group("/case")
	{
		c := casecnt.Controller{
			M: m,
		}

		caseRouter.POST("", c.CreateCase)
		caseRouter.GET("/active", c.GetActiveCases)
		caseRouter.GET("/unassigned", c.GetActiveCases)
		caseRouter.GET("/id/:caseID", c.GetCaseWithID)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
