package config

import (
	"BoxApi/errs"
	"encoding/json"
	"os"
)

type Config map[string]string

func NewConfig(namefile string) Config {
	var data = map[string]string{}

	f, err := os.ReadFile(namefile)

	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(f), &data)

	if data["Host"] == "" {
		panic(errs.ErrorNotFoundHost)
	}

	if data["Port"] == "" {
		panic(errs.ErrorNotFoundPort)
	}

	if data["PathJournal"] == "" {
		panic(errs.ErrorNotFoundPort)
	}

	return data
}
