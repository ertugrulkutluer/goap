package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string, env string) (string, error) {
	var path string
	// load .env file
	if env == "production" {
		path = "./environments/.env.production"
	} else if env == "staging" {
		path = "./environments/.env.staging"
	} else if env == "test" {
		path = "./environments/.env.test"
	} else {
		path = "./environments/.env.development"
	}
	err := godotenv.Load(path)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key), err
}
