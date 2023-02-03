package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config map[string]string

func NewConfig(namefile string) Config {
	f, err := os.ReadFile(namefile)
	if err != nil {
		log.Println(err)
	}
	var data = map[string]string{}
	json.Unmarshal([]byte(f), &data)
	fmt.Println(data)
	return data
}
