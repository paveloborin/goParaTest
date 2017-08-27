package helper

import (
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	PhpPath      string `json:"phpPath"`
	PhpUnitPath  string `json:"phpUnitPath"`
	TestDir      string `json:"dir"`
	ProcessesNum int `json:"processes"`
}

var config Config

func LoadConfiguration() Config {

	file := "./config.json"
	configFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
