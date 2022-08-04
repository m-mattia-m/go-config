package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var configData *Config
var singleton bool

func config() (*Config, int) {

	if singleton {
		return nil, 409 // already loaded
	}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("[CONFIG] : Error reading config file: " + err.Error())
	}
	fmt.Println("[CONFIG] : Successfully reading config file")

	err = yaml.Unmarshal(yamlFile, &configData)
	if err != nil {
		fmt.Println("[CONFIG] : Error to unmarshal config: " + err.Error())
	}
	fmt.Println("[CONFIG] : Successfully to unmarshal config")

	singleton = true
	return configData, 200
}

func GetDatabasePassword() string {
	config()
	fmt.Println("[CONFIG] : Successfully getting database password")
	return string(configData.Database.Password)
}

func GetDatabaseUsername() string {
	config()
	fmt.Println("[CONFIG] : Successfully getting database username")
	return string(configData.Database.Username)
}
