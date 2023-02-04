package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormCalc struct {
	X    float64 `form:"x" json:"x"`
	Y    float64 `form:"y" json:"y"`
	Func string  `form:"func" json:"func"`
}

func (s *Server) HandleCalcPost(c *gin.Context) {

	var json *FormCalc
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := s.Calc(*json)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (s *Server) HandleCalcGet(c *gin.Context) {
	page, err := ioutil.ReadFile("assets/calc.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, ContentTypeHTML, page)
}

func (s *Server) Calc(json FormCalc) string {
	var res float64
	switch json.Func {
	case "+":
		res = json.X + json.Y
	case "-":
		res = json.X - json.Y
	case "/":
		if json.Y == 0 {
			return fmt.Sprintf("Warning: 'What number multiplicate on zero is equal %v ?' Division by zero not allowed", json.X)
		}
		res = json.X / json.Y
	case "*":
		res = json.X * json.Y
	default:
		return "Support func: [`+` `-` `/` `*`]"
	}
	return fmt.Sprintf("%v", res)
}
