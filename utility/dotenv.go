package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string, env string) (string, error) {
	var path string
	// load .env file
	if path = "./environments/.env.production"; env == "staging" {
		path = "./environments/.env.staging"
	}
	err := godotenv.Load(path)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key), err
}
