package utils

import (
	"io/ioutil"
	"encoding/json"
	"project1/service/logger"
	"os"
)

var (
	//defines the app server port
	ServerPort string
	//defines Log level
	LogLevel string
)

func LoadEnvironmentConfiurations(filename string) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		logger.Errorf("File Error: %v/n", e)
	}
	ibusConfig := struct {
		ServerPort			string  `json:"SERVER_PORT"`
		LogLevel			string   `json:"LOG_LEVEL"`
	}{}
	json.Unmarshal(file, &ibusConfig)

	ServerPort = GetEnv("SERVER_PORT", ibusConfig.ServerPort)
	LogLevel = GetEnv("LOG_LEVEL", ibusConfig.LogLevel)
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	} 
	return value
}