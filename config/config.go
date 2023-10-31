package config

import (
	"log"
	"os"
)

// database connection configuration
type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

// retrieves the database connection configuration from environment variables
func GetConfig() Configuration {

	configuration := Configuration{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	log.Print(configuration)

	return configuration
}
