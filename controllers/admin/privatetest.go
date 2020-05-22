package admin

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (a Controller) Test(c *gin.Context) {

	officers, _ := a.M.GetUnassignedOfficers()

	//a.M.AssignCaseToOfficer("test")

	log.Println(officers)

	c.JSON(http.StatusOK, gin.H{
		"message": "secret",
		"officer": officers,
	})
}
