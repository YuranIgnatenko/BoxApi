package server

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleTimeGet(c *gin.Context) {
	page, err := ioutil.ReadFile("assets/time.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, ContentTypeHTML, page)
}

func (s *Server) HandleTimePost(c *gin.Context) {
	t := time.Now()
	c.JSON(http.StatusOK, gin.H{"data": t.String()})
}
