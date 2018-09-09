package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetConfiguration(environment string) map[string]string {
	configs := GetConfigurationFile(environment)
	if os.Getenv("MONGO_URL") != "" {
		configs["mongo"] = os.Getenv("MONGO_URL")
	}
	return configs
}

func GetConfigurationFile(environment string) map[string]string {
	directory, _ := os.Getwd()
	fmt.Printf("%s/config/%s.json", directory, environment)
	configBytes, _ := ioutil.ReadFile(fmt.Sprintf("%s/config/%s.json", directory, environment))
	configJSON := map[string]string{}
	json.Unmarshal(configBytes, &configJSON)
	return configJSON
}
