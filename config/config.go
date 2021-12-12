package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Represents database server and credentials
type configuration struct {
	Port       string
	Database   string
	JWT_Secret []byte
	Env        string
}

var Config = new(configuration)

// Read and parse the configuration file
func (c *configuration) Read(env string) *configuration {
	var path string
	// load .env file
	if env == "production" {
		path = "./environments/.env.production"
	} else if env == "staging" {
		path = "./environments/.env.staging"
	} else if env == "test" {
		path = "./environments/.env.test"
	} else {
		path = "./environments/.env"
	}
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	c.Database = os.Getenv("MONGO_URI")
	c.Port = os.Getenv("PORT")
	c.JWT_Secret = []byte(os.Getenv("JWT_SECRET"))
	c.Env = env
	return c
}
