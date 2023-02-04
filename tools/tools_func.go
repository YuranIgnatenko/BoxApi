package tools

import (
	"fmt"
	"strings"
)

func WarnString(any ...interface{}) string {
	result := ""
	for _, val := range any {
		result += fmt.Sprintf(" %v ", val)
	}
	return strings.TrimSpace(result)
}

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
			return WarnString(warningDivisionByZero, json.X)
		}
		res = json.X / json.Y
	case "*":
		res = json.X * json.Y
	default:
		return WarnString(labelSupport, supportedCalcFuncs)
	}
	return fmt.Sprintf("%v", res)
}

/*
Support params:

	cw - count all words
	cwu - count words upper-register (HELLO)
	cwt - count words title-register (Hello)
	cwl - count words lower-register (hello)
	cs -  count special symbols ("abc, abc" --> 1)
	lens - len all symbols ("abc, abc" --> 7)
	verb - verbose re-translate text ('Hello , Bob ?' --> 'Hello, Bob?')
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
		case "cwt":
			result[val] = cwt(json.Text)
		case "cs":
			result[val] = cs(json.Text)
		case "lens":
			result[val] = lens(json.Text)
		case "verb":
			result[verb(json.Text)] = 1
		default:
			result[fmt.Sprintf("Error: '%s'. %s: '%s' ", val, labelSupport, supportedWordPlusParams)] = -1
		}
	}

	return result
}

// split string ("Hello , Bob !" --> ["Hello,", "Bob!"])
func toSlice(text string) []string {
	var data = strings.Split(text, " ")
	var new_data = make([]string, 0)
	var buffer_ex_sym = ""
	var new_len_data = 0

	for _, word := range data {
		new_len_data += 1
		possible_ex_sym := strings.TrimSpace(word)
		if len(possible_ex_sym) > 1 {
			new_data = append(new_data, word)
			continue
		}
		for _, ex_sym := range exceptionSymbols {
			if possible_ex_sym == string(ex_sym) {
				buffer_ex_sym = possible_ex_sym
				new_len_data -= 1
				new_data[new_len_data-1] += buffer_ex_sym
				break
			}
		}
	}
	return new_data
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

// count words title-register
func cwt(text string) int {
	data := toSlice(text)
	count := 0
	for _, word := range data {
		if len(word) <= 1 {
			continue
		}
		s1, s2 := strings.ToUpper(word)[0], strings.ToLower(word)[1]
		w1, w2 := word[0], word[1]
		if w1 == s1 && w2 == s2 {
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

// count all symbols (after re-traslate text-!orginal text!) + space
func lens(text string) int {
	if strings.TrimSpace(text) == "" {
		return 0
	}
	return len(text)
}

// verbose re-translate text in code
func verb(text string) string {
	return strings.Join(toSlice(text), " ")
}

// count special symmbols ("#^&*()$")
func cs(text string) int {
	data := toSlice(text)
	count := 0
	for _, val := range data {
		for _, ex_sym := range exceptionSymbols {
			if string(val[len(val)-1]) == string(ex_sym) {
				count += 1
			}
		}
	}
	return count
}
