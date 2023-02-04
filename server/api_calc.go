package server

import (
	"io/ioutil"
	"net/http"

	"BoxApi/tools"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleCalcPost(c *gin.Context) {
	var json *tools.FormCalc
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := tools.Calc(*json)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (s *Server) HandleCalcGet(c *gin.Context) {
	page, err := ioutil.ReadFile("assets/calc.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, ContentTypeHTML, page)
}
