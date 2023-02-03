package config

import (
	"encoding/json"
	"os"
)

type Config map[string]string

func NewConfig(namefile string) Config {
	f, err := os.ReadFile(namefile)
	if err != nil {
		panic(err)
	}
	var data = map[string]string{}
	json.Unmarshal([]byte(f), &data)
	return data
}
