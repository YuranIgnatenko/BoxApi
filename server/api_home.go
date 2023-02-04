package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleHomeGet(c *gin.Context) {
	page, err := ioutil.ReadFile("assets/home.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, ContentTypeHTML, page)
}

func (s *Server) HandleHomePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "test"})
}
