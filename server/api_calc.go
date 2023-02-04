package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleCalcPost(c *gin.Context) {
	r := c.Request.PostForm
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("%#v", r)})
}

func (s *Server) HandleCalcGet(c *gin.Context) {
	page, err := ioutil.ReadFile("assets/calc.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, ContentTypeHTML, page)
}
