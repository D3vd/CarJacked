package casecnt

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ca Controller) CreateCase(c *gin.Context) {
	log.Println("success")
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}
