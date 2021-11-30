package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) (string, error) {

	// load .env file
	err := godotenv.Load("./environments/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key), err
}
