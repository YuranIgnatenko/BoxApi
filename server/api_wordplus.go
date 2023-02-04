package server

import (
	"net/http"

	"BoxApi/tools"

	"github.com/gin-gonic/gin"
)

// TODO: append html-page
func (s *Server) HandleWordPlusGet(c *gin.Context) {
}

func (s *Server) HandleWordPlusPost(c *gin.Context) {
	var json *tools.FormWordPlus
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := tools.LaunchWordPlus(*json)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
