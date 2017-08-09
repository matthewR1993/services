package util

import (
	"encoding/json"
	"os"
	"log"
)

type configuration struct {
	PgUsername, PgPassword string
}

// global variable to provide configuration
var AppConf configuration

func InitConfig(confPath string) {
	file, err := os.Open(confPath)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConf = configuration{}
	err = decoder.Decode(&AppConf)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}
}

