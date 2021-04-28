package config

import (
	"log"
	"path/filepath"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	absPath, err := filepath.Abs("../golang-training-university/config/home_config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = gonfig.GetConf(absPath, &configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}
