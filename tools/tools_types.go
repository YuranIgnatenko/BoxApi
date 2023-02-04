package tools

type FormCalc struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Func string  `json:"func"`
}

type FormWordPlus struct {
	Param []string `json:"params"`
	Text  string   `json:"text"`
}
