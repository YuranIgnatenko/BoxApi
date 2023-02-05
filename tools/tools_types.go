package tools

type FormCalc struct {
	X    float64 `form:"x" json:"X" binding:"required"`
	Y    float64 `form:"y" json:"Y" binding:"required"`
	Func string  `form="func" json:"Func" binding:"required"`
}

type FormWordPlus struct {
	Param []string `json:"params"`
	Text  string   `json:"text"`
}
