package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PayoutConfiguration struct {
	TravelApiHost string
	ApiKey        string
	AccessToken   string
}

type PostgresConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func LoadBasiDataScrapperConfig() (configuration PayoutConfiguration) {
	fileName, _ := filepath.Abs("../travel-planner/config.json")
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = PayoutConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func LoadDatabaseConfig() (configuration PostgresConfiguration) {
	fileName, _ := filepath.Abs("../travel-planner/config.json")
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = PostgresConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
