package tools

import (
	"fmt"
	"strings"
)

/*
Support funcs:
  - add
  - sub
  - mul
    / div
*/
func Calc(json FormCalc) string {
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

/*
Support params:

	cw - count all words
	cwu - count words upper-register
	cwl - count words lower-register
	cs - len text (count all symbols)
*/
func LaunchWordPlus(json FormWordPlus) map[string]int {
	result := make(map[string]int)
	for _, val := range json.Param {
		switch val {
		case "cw":
			result[val] = cw(json.Text)
		case "cwu":
			result[val] = cwu(json.Text)
		case "cwl":
			result[val] = cwl(json.Text)
		case "cs":
			result[val] = cs(json.Text)
		default:
			result[fmt.Sprintf("Warning! Error param: '%s' supports:[cw,cwu,cwl,cs]", val)] = -1
		}
	}

	return result
}

// TODO: danger-situate: "hello , Bob" --> ['hello', ',', 'Bob']
// add ignore [. , + / ? ! @ ....]
func toSlice(text string) []string {
	return strings.Split(text, " ")
}

// count all words
func cw(text string) int {
	data := toSlice(text)
	return len(data)
}

// count words upper-register
func cwu(text string) int {
	data := toSlice(text)
	count := 0
	for _, word := range data {
		if strings.ToUpper(word) == word {
			count += 1
		}
	}
	return count
}

// count words lower-register
func cwl(text string) int {
	data := toSlice(text)
	count := 0
	for _, word := range data {
		if strings.ToLower(word) == word {
			count += 1
		}
	}
	return count
}

// count all symbols
func cs(text string) int {
	if strings.TrimSpace(text) == "" {
		return 0
	}
	c := 0
	for ind, _ := range text {
		c += ind
	}
	return c
}
