package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a Controller) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "secret",
	})
}
