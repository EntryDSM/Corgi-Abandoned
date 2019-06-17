package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Registry             string
	User                 string
	Password             string
	PollingRatePerMinute int
	ServerCertFilePath   string
}

var CorgiConfig Config

func InitializeConfig(configFilePath string) {
	jsonFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(jsonFile, &CorgiConfig)

	if err != nil {
		panic(err.Error())
	}
}
