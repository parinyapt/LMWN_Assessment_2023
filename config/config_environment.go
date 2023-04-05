package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvironmentFileSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error]->Failed to load environment file : %s", err)
	}
}

func EnvironmentVariableCheck() {
	var requireEnvVariableList = []string{
		"PORT",
		"TZ",
	}

	for _, v := range requireEnvVariableList {
		if len([]byte(os.Getenv(v))) == 0 {
			log.Fatalf("[Error]->Environment Variable '%s' is not set", v)
		}
	}
}