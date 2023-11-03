package conf

import (
	"ace/model"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func GetConf(path string) model.Config {
	var config model.Config

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Cannot read config file, failed path, error: %v \n", err.Error())
		os.Exit(0)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Cannot deal with config data, error: %v \n", err.Error())
	}

	return config
}
