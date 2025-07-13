package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()
	if err != nil {
		// Log the error if the .env file not found or cannot be loaded
		fmt.Println("Error loading .env file")
	}
}

func GetString(key string, fallback string) string {
	load()

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}
