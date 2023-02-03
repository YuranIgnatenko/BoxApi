package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

type Server struct {
	Host string
	Port string
}

func NewServer(host, port string) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s *Server) Addr() string {
	return s.Host + ":" + s.Port
}

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
